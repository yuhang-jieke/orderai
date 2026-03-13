# 秒杀系统数据库 ER 图

## Mermaid ER 图

```mermaid
erDiagram
    PRODUCTS ||--o{ FLASH_SALES : "1:N"
    FLASH_SALES ||--o{ SECKILL_ORDERS : "1:N"
    FLASH_SALES ||--o{ SECKILL_STOCK_LOCK : "1:N"
    FLASH_SALES ||--o{ SECKILL_LOGS : "1:N"
    FLASH_SALES ||--o{ USER_ACTIVITY : "1:N"
    FLASH_SALES ||--o{ FLASH_SALE_INVENTORY : "1:N"
    SECKILL_ORDERS ||--o{ SECKILL_ORDER_ITEMS : "1:N"
    PRODUCTS ||--o{ SECKILL_ORDER_ITEMS : "1:N"
    
    PRODUCTS {
        bigint id PK "商品ID"
        varchar sku UK "SKU编码"
        varchar name "商品名称"
        text description "商品描述"
        bigint category_id "分类ID"
        decimal price "原价"
        int original_stock "原始库存"
        int current_stock "当前库存"
        tinyint status "状态"
        datetime created_at "创建时间"
        datetime updated_at "更新时间"
    }
    
    FLASH_SALES {
        bigint id PK "活动ID"
        varchar title "活动标题"
        bigint product_id FK "商品ID"
        datetime start_time "开始时间"
        datetime end_time "结束时间"
        decimal seckill_price "秒杀价"
        int total_quantity "总数量"
        int remaining_quantity "剩余数量"
        int max_per_user "每人限购"
        tinyint status "状态"
        datetime created_at "创建时间"
        datetime updated_at "更新时间"
    }
    
    SECKILL_ORDERS {
        bigint id PK "订单ID"
        varchar order_sn UK "订单编号"
        bigint flash_sale_id FK "活动ID"
        bigint product_id FK "商品ID"
        bigint user_id "用户ID"
        int quantity "购买数量"
        decimal total_amount "订单金额"
        tinyint status "订单状态"
        datetime pay_time "支付时间"
        datetime cancel_time "取消时间"
        datetime finish_time "完成时间"
        datetime created_at "创建时间"
        datetime updated_at "更新时间"
    }
    
    SECKILL_STOCK_LOCK {
        bigint id PK "锁定ID"
        bigint flash_sale_id FK "活动ID"
        bigint user_id "用户ID"
        varchar order_sn "订单SN"
        int quantity "锁定数量"
        datetime lock_time "锁定时间"
        datetime expire_time "过期时间"
        tinyint status "状态"
        datetime created_at "创建时间"
    }
    
    SECKILL_LOGS {
        bigint id PK "日志ID"
        bigint flash_sale_id FK "活动ID"
        bigint user_id "用户ID"
        tinyint action_type "操作类型"
        tinyint action_result "操作结果"
        varchar request_ip "请求IP"
        varchar user_agent "User-Agent"
        json details "详细信息"
        datetime created_at "创建时间"
    }
    
    USER_ACTIVITY {
        bigint id PK "记录ID"
        bigint user_id "用户ID"
        bigint flash_sale_id FK "活动ID"
        tinyint action_type "操作类型"
        datetime request_time "请求时间"
        varchar request_ip "请求IP"
        datetime created_at "创建时间"
    }
    
    SECKILL_ORDER_ITEMS {
        bigint id PK "明细ID"
        bigint order_id FK "订单ID"
        bigint product_id FK "商品ID"
        varchar product_name "商品名称"
        varchar product_sku "商品SKU"
        int quantity "数量"
        decimal seckill_price "秒杀单价"
        decimal total_price "小计"
        datetime created_at "创建时间"
    }
    
    FLASH_SALE_INVENTORY {
        bigint id PK "库存ID"
        bigint flash_sale_id FK "活动ID"
        bigint warehouse_id "仓库ID"
        int allocated_quantity "分配数量"
        int available_quantity "可用数量"
        datetime created_at "创建时间"
        datetime updated_at "更新时间"
    }
```

---

## 表关系说明

### 核心业务流程关系

```
┌─────────────┐      ┌─────────────────┐      ┌─────────────────┐
│  products   │◄─────┤  flash_sales    │◄─────┤ seckill_orders  │
│   (商品)    │ 1:N  │   (秒杀活动)     │ 1:N  │   (秒杀订单)     │
└─────────────┘      └─────────────────┘      └─────────────────┘
                              │
              ┌───────────────┼───────────────┐
              │               │               │
              ▼               ▼               ▼
┌──────────────────┐ ┌──────────────┐ ┌────────────────┐
│seckill_stock_lock│ │ seckill_logs │ │  user_activity │
│   (库存锁定)      │ │  (操作日志)   │ │  (行为追踪)     │
└──────────────────┘ └──────────────┘ └────────────────┘
```

---

## 外键关系明细

| 主表 | 子表 | 外键字段 | 关系类型 | 说明 |
|------|------|----------|----------|------|
| **products** | flash_sales | product_id | 1:N | 一个商品可参与多次秒杀 |
| **products** | seckill_orders | product_id | 1:N | 一个商品可产生多个订单 |
| **products** | seckill_order_items | product_id | 1:N | 一个商品可出现在多个订单项 |
| **flash_sales** | seckill_orders | flash_sale_id | 1:N | 一个活动可产生多个订单 |
| **flash_sales** | seckill_stock_lock | flash_sale_id | 1:N | 一个活动可有多个库存锁定 |
| **flash_sales** | seckill_logs | flash_sale_id | 1:N | 一个活动产生多条日志 |
| **flash_sales** | user_activity | flash_sale_id | 1:N | 一个活动有多个用户行为 |
| **flash_sales** | flash_sale_inventory | flash_sale_id | 1:N | 一个活动分布在多个仓库 |
| **seckill_orders** | seckill_order_items | order_id | 1:N | 一个订单包含多个商品项 |

---

## 唯一约束

| 表名 | 唯一约束 | 字段 | 说明 |
|------|----------|------|------|
| products | uk_sku | sku | SKU唯一 |
| seckill_orders | uk_order_sn | order_sn | 订单号唯一 |
| seckill_stock_lock | uk_user_flash | user_id + flash_sale_id | **防止重复抢购** |
| flash_sale_inventory | uk_flash_warehouse | flash_sale_id + warehouse_id | 活动+仓库唯一 |

---

## PlantUML 版本

```plantuml
@startuml

skinparam classAttributeIconSize 0
skinparam linetype ortho
skinparam monochrome true

class products {
  + id: bigint <<PK>>
  + sku: varchar <<UK>>
  + name: varchar
  + description: text
  + category_id: bigint
  + price: decimal
  + original_stock: int
  + current_stock: int
  + status: tinyint
  + created_at: datetime
  + updated_at: datetime
}

class flash_sales {
  + id: bigint <<PK>>
  + title: varchar
  + product_id: bigint <<FK>>
  + start_time: datetime
  + end_time: datetime
  + seckill_price: decimal
  + total_quantity: int
  + remaining_quantity: int
  + max_per_user: int
  + status: tinyint
  + created_at: datetime
  + updated_at: datetime
}

class seckill_orders {
  + id: bigint <<PK>>
  + order_sn: varchar <<UK>>
  + flash_sale_id: bigint <<FK>>
  + product_id: bigint <<FK>>
  + user_id: bigint
  + quantity: int
  + total_amount: decimal
  + status: tinyint
  + pay_time: datetime
  + cancel_time: datetime
  + finish_time: datetime
  + created_at: datetime
  + updated_at: datetime
}

class seckill_stock_lock {
  + id: bigint <<PK>>
  + flash_sale_id: bigint <<FK>>
  + user_id: bigint
  + order_sn: varchar
  + quantity: int
  + lock_time: datetime
  + expire_time: datetime
  + status: tinyint
  + created_at: datetime
}

class seckill_logs {
  + id: bigint <<PK>>
  + flash_sale_id: bigint <<FK>>
  + user_id: bigint
  + action_type: tinyint
  + action_result: tinyint
  + request_ip: varchar
  + user_agent: varchar
  + details: json
  + created_at: datetime
}

class user_activity {
  + id: bigint <<PK>>
  + user_id: bigint
  + flash_sale_id: bigint <<FK>>
  + action_type: tinyint
  + request_time: datetime
  + request_ip: varchar
  + created_at: datetime
}

class seckill_order_items {
  + id: bigint <<PK>>
  + order_id: bigint <<FK>>
  + product_id: bigint <<FK>>
  + product_name: varchar
  + product_sku: varchar
  + quantity: int
  + seckill_price: decimal
  + total_price: decimal
  + created_at: datetime
}

class flash_sale_inventory {
  + id: bigint <<PK>>
  + flash_sale_id: bigint <<FK>>
  + warehouse_id: bigint
  + allocated_quantity: int
  + available_quantity: int
  + created_at: datetime
  + updated_at: datetime
}

products ||--o{ flash_sales : "1:N"
products ||--o{ seckill_orders : "1:N"
products ||--o{ seckill_order_items : "1:N"
flash_sales ||--o{ seckill_orders : "1:N"
flash_sales ||--o{ seckill_stock_lock : "1:N"
flash_sales ||--o{ seckill_logs : "1:N"
flash_sales ||--o{ user_activity : "1:N"
flash_sales ||--o{ flash_sale_inventory : "1:N"
seckill_orders ||--o{ seckill_order_items : "1:N"

note right of seckill_stock_lock
  **唯一约束**: uk_user_flash
  (user_id, flash_sale_id)
  防止同一用户重复抢购
end note

note bottom of flash_sales
  **触发器**:
  - after_insert: 扣减库存
  - after_update: 恢复库存
end note

@enduml
```

---

## 使用说明

### 1. Mermaid 渲染
在支持 Mermaid 的 Markdown 编辑器（如 VS Code + Mermaid 插件、GitHub）中直接查看。

### 2. PlantUML 渲染
- 在线工具：[plantuml.com/plantuml](https://plantuml.com/plantuml)
- VS Code 插件：PlantUML
- IntelliJ 插件：PlantUML integration

### 3. 导出图片
使用 PlantUML 可以导出 PNG/SVG/PDF 格式：
```bash
# 命令行导出
java -jar plantuml.jar er-diagram.puml -tpng
java -jar plantuml.jar er-diagram.puml -tsvg
```

---

## 关键设计亮点

| 设计点 | 实现方式 | 优势 |
|--------|----------|------|
| **防重复抢购** | seckill_stock_lock.uk_user_flash | 数据库唯一约束，原子性保证 |
| **库存一致性** | 触发器自动扣减/恢复 | 避免应用层遗漏 |
| **高并发查询** | 12+战略索引 | 覆盖所有热查询 |
| **审计追踪** | seckill_logs + user_activity | 完整操作链路 |
| **扩展性** | 可选的多仓库/多商品支持 | 未来业务扩展 |
