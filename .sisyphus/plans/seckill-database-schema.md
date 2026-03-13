# Flash Sale E-commerce Database Schema

## TL;DR

> **Quick Summary**: Complete MySQL 8.0 InnoDB database schema for e-commerce order system supporting flash sale activities. Includes 8 tables with full constraints, 12+ strategic indexes, 2 triggers, and Go struct mappings.
> 
> **Deliverables**:
> - 8 complete database tables (products, flash_sales, seckill_orders, stock_lock, logs, user_activity, inventory, order_items)
> - 12+ performance-optimized indexes
> - 2 inventory management triggers
> - Sample data and common queries
> - Go struct documentation
> 
> **Estimated Effort**: Small
> **Parallel Execution**: YES - 2 waves
> **Critical Path**: Tasks 1-4 → 5-7 → F1-F4

---

## Context

### Original Request
User requested: "设计一个支持秒杀活动的电商订单系统库表 (MySQL 8.0, InnoDB 引擎)"

### Interview Summary
**Key Discussions**:
- Requirements confirmed: "Basic flash sales" - simple flash sales with product quantity limits and time constraints
- Database: MySQL 8.0, InnoDB engine
- Project: Go-based backend (will generate Go structs from schema)

**Research Findings**:
- InnoDB supports row-level locking and MVCC for high concurrency
- Optimistic locking via stock_lock table prevents race conditions
- Basic approach: product table + flash_sale table + order table + stock management

### Metis Review
**Identified Gaps** (addressed):
- Concurrency control: Added seckill_stock_lock table with unique constraint per user+flash_sale
- Audit trail: Added seckill_logs table for all actions
- Tracking: Added user_activity table for behavior monitoring
- Performance: Added compound indexes for hot queries

---

## Work Objectives

### Core Objective
Design and document a production-ready MySQL schema supporting basic flash sale scenarios with efficient concurrency control and comprehensive audit logging.

### Concrete Deliverables
- Complete database schema with 8 tables
- 12+ strategic indexes for high-traffic queries
- 2 triggers for inventory management
- Sample data and common query patterns
- Go struct documentation

### Definition of Done
- [ ] All tables created with proper constraints
- [ ] All indexes built
- [ ] Both triggers active
- [ ] Schema compatible with Go code generation

### Must Have
- Optimistic locking (seckill_stock_lock with unique constraint)
- Inventory tracking (remaining_quantity field)
- Comprehensive audit logging (seckill_logs)
- Efficient query patterns (strategic indexes)
- Go type compatibility (BIGINT UNSIGNED → uint64)

### Must NOT Have (Guardrails)
- No distributed coordination
- No advanced anti-spam features
- No complex inventory distribution

---

## Verification Strategy

### Test Decision
- **Infrastructure exists**: NO
- **Automated tests**: None (schema design only)
- **Agent-Executed QA**: ALWAYS (mandatory for all tasks)

### QA Policy
Every task MUST include agent-executed QA scenarios.
Evidence saved to `.sisyphus/evidence/task-{N}-{scenario-slug}.{ext}`.

---

## Execution Strategy

### Parallel Execution Waves

```
Wave 1 (Start Immediately):
├── Task 1: Create main tables (products, flash_sales, seckill_orders)
├── Task 2: Create supporting tables (stock_lock, logs, user_activity)
├── Task 3: Create optional tables (inventory, order_items)
└── Task 4: Add strategic indexes for performance

Wave 2 (After Wave 1):
├── Task 5: Add triggers for inventory management
├── Task 6: Document schema with Go struct mappings
└── Task 7: Add sample data for testing

Wave FINAL (After ALL tasks):
├── Task F1: Plan compliance audit (oracle)
├── Task F2: Schema validation (deep)
├── Task F3: Query performance verification (deep)
└── Task F4: Scope fidelity check (deep)
```

---

## TODOs

- [ ] 1. Create main tables (products, flash_sales, seckill_orders)

  **What to do**:
  - Create products table with SKU, name, category, pricing, stock management
  - Create flash_sales table with time constraints and quantity limits
  - Create seckill_orders table with order lifecycle tracking

  **Recommend Agent Profile**: `quick` with `sql` skill

  **QA Scenarios**:
  \`\`\`
  Scenario: Main tables created successfully
    Tool: Bash (mysql client)
    Steps:
      1. Run CREATE TABLE statements
      2. Verify table names exist via SHOW TABLES
    Expected Result: All 3 tables created with expected columns
    Evidence: .sisyphus/evidence/task-1-validation.txt
  \`\`\`

  **Commit**: YES - `feat(schema): add main tables for flash sale system`

- [ ] 2. Create supporting tables (stock_lock, logs, user_activity)

  **What to do**:
  - Create seckill_stock_lock with unique constraint
  - Create seckill_logs for audit trail
  - Create user_activity for behavior tracking

  **Recommend Agent Profile**: `quick` with `sql` skill

  **Commit**: YES

- [ ] 3. Create optional tables (inventory, order_items)

  **What to do**:
  - Create flash_sale_inventory for multi-warehouse
  - Create seckill_order_items for multi-item orders

  **Recommend Agent Profile**: `quick` with `sql` skill

  **Commit**: YES

- [ ] 4. Add strategic indexes for performance

  **What to do**:
  - Add indexes for hot queries (active flash sales, product stock, orders)
  - Create compound indexes for common filter patterns

  **Recommend Agent Profile**: `quick` with `sql`, `performance` skills

  **QA Scenarios**:
  \`\`\`
  Scenario: Hot query uses correct index
    Tool: Bash (mysql EXPLAIN)
    Steps:
      1. Run EXPLAIN on active flash sales query
      2. Verify idx_active_sales used
    Expected Result: All hot queries use strategic indexes
    Evidence: .sisyphus/evidence/task-4-index-verify.txt
  \`\`\`

  **Commit**: YES

- [ ] 5. Add triggers for inventory management

  **What to do**:
  - Create trigger for order insert (reduce remaining_quantity)
  - Create trigger for order cancel (restore remaining_quantity)

  **Recommend Agent Profile**: `quick` with `sql` skill

  **QA Scenarios**:
  \`\`\`
  Scenario: Trigger reduces inventory on order
    Tool: Bash (mysql client)
    Preconditions: Flash sale with remaining_quantity = 10
    Steps:
      1. Insert seckill_order for quantity 2
      2. Query flash_sales.remaining_quantity
    Expected Result: remaining_quantity = 8
    Evidence: .sisyphus/evidence/task-5-trigger-test.txt
  \`\`\`

  **Commit**: YES

- [ ] 6. Document schema with Go struct mappings

  **What to do**:
  - Document each table with Go struct comments
  - Map column types to Go types (BIGINT UNSIGNED → uint64)

  **Recommend Agent Profile**: `deep` with `golang`, `writing`, `sql` skills

  **Commit**: YES

- [ ] 7. Add sample data for testing

  **What to do**:
  - Insert sample products, flash sales, orders, logs

  **Recommend Agent Profile**: `quick` with `sql` skill

  **Commit**: YES

---

## Final Verification Wave

- [ ] F1. **Plan Compliance Audit** — `oracle`
- [ ] F2. **Schema Validation** — `deep`
- [ ] F3. **Query Performance Verification** — `deep`
- [ ] F4. **Scope Fidelity Check** — `deep`

---

## Commit Strategy

- **1-7**: `type(schema): create flash sale system tables` — seckill-database-schema.md

---

## Success Criteria

### Verification Commands
```bash
mysql < seckill-database-schema.md
mysql -e "SHOW TABLES;"
mysql -e "SELECT TABLE_NAME, INDEX_NAME FROM information_schema.STATISTICS WHERE TABLE_SCHEMA = 'seckill_demo';"
```

### Final Checklist
- [ ] All "Must Have" present
- [ ] All "Must NOT Have" absent
- [ ] All indexes created and tested
- [ ] All triggers active and verified
- [ ] Go struct mappings complete

---

## Complete Database Schema

### 1. Products (商品表)
```sql
CREATE TABLE products (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品ID',
    sku VARCHAR(64) NOT NULL COMMENT '商品SKU编码',
    name VARCHAR(255) NOT NULL COMMENT '商品名称',
    description TEXT COMMENT '商品描述',
    category_id BIGINT UNSIGNED NOT NULL COMMENT '分类ID',
    price DECIMAL(10,2) NOT NULL COMMENT '原价',
    original_stock INT NOT NULL DEFAULT 0 COMMENT '原始库存数量',
    current_stock INT NOT NULL DEFAULT 0 COMMENT '当前可用库存',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '状态: 0=下架, 1=上架',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    PRIMARY KEY (id),
    UNIQUE KEY uk_sku (sku),
    KEY idx_category (category_id),
    KEY idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品表';
```

### 2. Flash Sales (秒杀活动表)
```sql
CREATE TABLE flash_sales (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '秒杀活动ID',
    title VARCHAR(255) NOT NULL COMMENT '活动标题',
    product_id BIGINT UNSIGNED NOT NULL COMMENT '商品ID',
    start_time DATETIME NOT NULL COMMENT '开始时间',
    end_time DATETIME NOT NULL COMMENT '结束时间',
    seckill_price DECIMAL(10,2) NOT NULL COMMENT '秒杀价',
    total_quantity INT NOT NULL DEFAULT 0 COMMENT '秒杀总数量',
    remaining_quantity INT NOT NULL DEFAULT 0 COMMENT '剩余可秒杀数量',
    max_per_user INT NOT NULL DEFAULT 1 COMMENT '每人最多可购买数量',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '状态: 0=未开始, 1=进行中, 2=已结束, 3=已取消',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    PRIMARY KEY (id),
    KEY idx_product_time (product_id, start_time, end_time),
    KEY idx_status_time (status, start_time),
    KEY idx_end_time (end_time)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='秒杀活动表';
```

### 3. Seckill Orders (秒杀订单表)
```sql
CREATE TABLE seckill_orders (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单ID',
    order_sn VARCHAR(32) NOT NULL COMMENT '订单编号',
    flash_sale_id BIGINT UNSIGNED NOT NULL COMMENT '秒杀活动ID',
    product_id BIGINT UNSIGNED NOT NULL COMMENT '商品ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    quantity INT NOT NULL DEFAULT 1 COMMENT '购买数量',
    total_amount DECIMAL(10,2) NOT NULL COMMENT '订单总金额',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '订单状态: 0=待支付, 1=已支付, 2=已取消, 3=已退款, 4=已完成',
    pay_time DATETIME(3) DEFAULT NULL COMMENT '支付时间',
    cancel_time DATETIME(3) DEFAULT NULL COMMENT '取消时间',
    finish_time DATETIME(3) DEFAULT NULL COMMENT '完成时间',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    PRIMARY KEY (id),
    UNIQUE KEY uk_order_sn (order_sn),
    KEY idx_user (user_id),
    KEY idx_flash_sale (flash_sale_id),
    KEY idx_status_created (status, created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='秒杀订单表';
```

### 4. Seckill Stock Lock (库存锁定表)
```sql
CREATE TABLE seckill_stock_lock (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '库存锁定ID',
    flash_sale_id BIGINT UNSIGNED NOT NULL COMMENT '秒杀活动ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    order_sn VARCHAR(32) NOT NULL COMMENT '关联的订单SN',
    quantity INT NOT NULL DEFAULT 0 COMMENT '锁定数量',
    lock_time DATETIME(3) NOT NULL COMMENT '锁定时间',
    expire_time DATETIME(3) NOT NULL COMMENT '过期时间',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '状态: 1=锁定中, 2=已释放, 3=已扣除',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    PRIMARY KEY (id),
    UNIQUE KEY uk_user_flash (user_id, flash_sale_id),
    KEY idx_expire_time (expire_time),
    KEY idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='秒杀库存锁定表';
```

### 5. Seckill Logs (秒杀日志表)
```sql
CREATE TABLE seckill_logs (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志ID',
    flash_sale_id BIGINT UNSIGNED NOT NULL COMMENT '秒杀活动ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    action_type TINYINT NOT NULL COMMENT '操作类型: 1=抢购尝试, 2=抢购成功, 3=抢购失败, 4=订单创建, 5=支付成功',
    action_result TINYINT NOT NULL DEFAULT 0 COMMENT '结果: 0=未知, 1=成功, 2=库存不足, 3=活动未开始, 4=活动已结束, 5=限购超限, 6=重复抢购',
    request_ip VARCHAR(45) COMMENT '请求IP',
    user_agent VARCHAR(512) COMMENT 'User-Agent',
    details JSON COMMENT '详细信息 (JSON格式)',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    PRIMARY KEY (id),
    KEY idx_flash_sale_action (flash_sale_id, action_type),
    KEY idx_user_time (user_id, created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='秒杀活动日志表';
```

### 6. User Activity (用户行为追踪表)
```sql
CREATE TABLE user_activity (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    flash_sale_id BIGINT UNSIGNED NOT NULL COMMENT '秒杀活动ID',
    action_type TINYINT NOT NULL COMMENT '操作类型: 1=进入页面, 2=点击抢购, 3=提交订单',
    request_time DATETIME(3) NOT NULL COMMENT '请求时间',
    request_ip VARCHAR(45) COMMENT '请求IP',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    PRIMARY KEY (id),
    KEY idx_user_flash_time (user_id, flash_sale_id, request_time)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户行为追踪表';
```

### 7. Flash Sale Inventory (秒杀库存分发表 - 可选)
```sql
CREATE TABLE flash_sale_inventory (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '库存记录ID',
    flash_sale_id BIGINT UNSIGNED NOT NULL COMMENT '秒杀活动ID',
    warehouse_id BIGINT UNSIGNED NOT NULL COMMENT '仓库ID',
    allocated_quantity INT NOT NULL DEFAULT 0 COMMENT '分配数量',
    available_quantity INT NOT NULL DEFAULT 0 COMMENT '可用数量',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    PRIMARY KEY (id),
    UNIQUE KEY uk_flash_warehouse (flash_sale_id, warehouse_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='秒杀库存分发表';
```

### 8. Seckill Order Items (秒杀订单明细表 - 可选)
```sql
CREATE TABLE seckill_order_items (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单明细ID',
    order_id BIGINT UNSIGNED NOT NULL COMMENT '订单ID',
    product_id BIGINT UNSIGNED NOT NULL COMMENT '商品ID',
    product_name VARCHAR(255) NOT NULL COMMENT '商品名称',
    product_sku VARCHAR(64) NOT NULL COMMENT '商品SKU',
    quantity INT NOT NULL DEFAULT 1 COMMENT '购买数量',
    seckill_price DECIMAL(10,2) NOT NULL COMMENT '秒杀单价',
    total_price DECIMAL(10,2) NOT NULL COMMENT '小计',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    PRIMARY KEY (id),
    KEY idx_order (order_id),
    KEY idx_product (product_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='秒杀订单明细表';
```

---

## Strategic Indexes (12+ indexes)

```sql
-- Maintenance auto-generated unique index (not a separate statement, part of table definition)
-- uk_user_flash on seckill_stock_lock: (user_id, flash_sale_id)

-- Additional strategic indexes

-- Active flash sales query
CREATE INDEX idx_active_sales ON flash_sales (status, start_time, end_time);

-- Product availability check
CREATE INDEX idx_product_stock ON products (id, current_stock);

-- Order user query
CREATE INDEX idx_order_user_status ON seckill_orders (user_id, status, created_at);

-- Concurrent order creation
CREATE INDEX idx_seckill_user ON seckill_orders (flash_sale_id, user_id);
```

---

## Triggers

### Trigger 1: Order Insert - Reduce Inventory
```sql
DELIMITER $$
CREATE TRIGGER after_seckill_order_insert
AFTER INSERT ON seckill_orders
FOR EACH ROW
BEGIN
    UPDATE flash_sales 
    SET remaining_quantity = remaining_quantity - NEW.quantity
    WHERE id = NEW.flash_sale_id;
END$$
DELIMITER ;
```

### Trigger 2: Order Cancel - Restore Inventory
```sql
DELIMITER $$
CREATE TRIGGER after_seckill_order_cancel
AFTER UPDATE ON seckill_orders
FOR EACH ROW
BEGIN
    IF OLD.status IN (0, 1) AND NEW.status = 2 THEN
        UPDATE flash_sales 
        SET remaining_quantity = remaining_quantity + OLD.quantity
        WHERE id = OLD.flash_sale_id;
    END IF;
END$$
DELIMITER ;
```

---

## Go Struct Mappings

### Product
```go
type Product struct {
    ID            uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
    SKU           string     `gorm:"uniqueIndex;not null" json:"sku"`
    Name          string     `gorm:"not null" json:"name"`
    Description   *string    `gorm:"type:text" json:"description,omitempty"`
    CategoryID    uint64     `gorm:"not null" json:"category_id"`
    Price         float64    `gorm:"type:decimal(10,2);not null" json:"price"`
    OriginalStock int        `gorm:"not null;default:0" json:"original_stock"`
    CurrentStock  int        `gorm:"not null;default:0" json:"current_stock"`
    Status        int        `gorm:"not null;default:1" json:"status"` // 0=下架, 1=上架
    CreatedAt     time.Time  `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`
    UpdatedAt     time.Time  `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);autoUpdate" json:"updated_at"`
}
```

### FlashSale
```go
type FlashSale struct {
    ID                uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
    Title             string     `gorm:"not null" json:"title"`
    ProductID         uint64     `gorm:"not null" json:"product_id"`
    StartTime         time.Time  `gorm:"not null" json:"start_time"`
    EndTime           time.Time  `gorm:"not null" json:"end_time"`
    SeckillPrice      float64    `gorm:"type:decimal(10,2);not null" json:"seckill_price"`
    TotalQuantity     int        `gorm:"not null;default:0" json:"total_quantity"`
    RemainingQuantity int        `gorm:"not null;default:0" json:"remaining_quantity"`
    MaxPerUser        int        `gorm:"not null;default:1" json:"max_per_user"`
    Status            int        `gorm:"not null;default:0" json:"status"` // 0=未开始, 1=进行中, 2=已结束, 3=已取消
    CreatedAt         time.Time  `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`
    UpdatedAt         time.Time  `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);autoUpdate" json:"updated_at"`
}
```

### SeckillOrder
```go
type SeckillOrder struct {
    ID           uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
    OrderSN      string      `gorm:"uniqueIndex;not null" json:"order_sn"`
    FlashSaleID  uint64      `gorm:"not null" json:"flash_sale_id"`
    ProductID    uint64      `gorm:"not null" json:"product_id"`
    UserID       uint64      `gorm:"not null" json:"user_id"`
    Quantity     int         `gorm:"not null;default:1" json:"quantity"`
    TotalAmount  float64     `gorm:"type:decimal(10,2);not null" json:"total_amount"`
    Status       int         `gorm:"not null;default:0" json:"status"` // 0=待支付, 1=已支付, 2=已取消, 3=已退款, 4=已完成
    PayTime      *time.Time  `gorm:"type:datetime(3)" json:"pay_time,omitempty"`
    CancelTime   *time.Time  `gorm:"type:datetime(3)" json:"cancel_time,omitempty"`
    FinishTime   *time.Time  `gorm:"type:datetime(3)" json:"finish_time,omitempty"`
    CreatedAt    time.Time   `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`
    UpdatedAt    time.Time   `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);autoUpdate" json:"updated_at"`
}
```

### StockLock
```go
type StockLock struct {
    ID         uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
    FlashSaleID uint64    `gorm:"not null" json:"flash_sale_id"`
    UserID     uint64     `gorm:"not null" json:"user_id"`
    OrderSN    string     `gorm:"not null" json:"order_sn"`
    Quantity   int        `gorm:"not null;default:0" json:"quantity"`
    LockTime   time.Time  `gorm:"not null" json:"lock_time"`
    ExpireTime time.Time  `gorm:"not null" json:"expire_time"`
    Status     int        `gorm:"not null;default:1" json:"status"` // 1=锁定中, 2=已释放, 3=已扣除
    CreatedAt  time.Time  `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`
}
// Unique constraint on (user_id, flash_sale_id) - prevents duplicate抢购
```

### SeckillLog
```go
type SeckillLog struct {
    ID          uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
    FlashSaleID uint64      `gorm:"not null" json:"flash_sale_id"`
    UserID      uint64      `gorm:"not null" json:"user_id"`
    ActionType  int         `gorm:"not null" json:"action_type"`
    ActionResult int        `gorm:"not null;default:0" json:"action_result"`
    RequestIP   *string     `gorm:"type:varchar(45)" json:"request_ip,omitempty"`
    UserAgent   *string     `gorm:"type:varchar(512)" json:"user_agent,omitempty"`
    Details     *string     `gorm:"type:json" json:"details,omitempty"`
    CreatedAt   time.Time   `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`
}
```

---

## Sample Data

```sql
-- Insert sample products
INSERT INTO products (sku, name, category_id, price, original_stock, current_stock, status) 
VALUES 
('PROD001', 'iPhone 15 Pro', 1, 8999.00, 100, 100, 1),
('PROD002', 'Samsung Galaxy S24', 1, 6999.00, 50, 50, 1),
('PROD003', 'Xiaomi 14', 1, 3999.00, 200, 200, 1),
('PROD004', 'Huawei Mate 60', 1, 5999.00, 80, 80, 1),
('PROD005', 'OPPO Find X6', 1, 4999.00, 120, 120, 1);

-- Insert sample flash sales
INSERT INTO flash_sales (title, product_id, start_time, end_time, seckill_price, total_quantity, remaining_quantity, max_per_user, status) 
VALUES 
('iPhone 15 Pro 限时秒杀', 1, '2026-03-15 10:00:00', '2026-03-15 10:30:00', 7999.00, 10, 10, 1, 0),
('Samsung Galaxy S24 抢购', 2, '2026-03-15 12:00:00', '2026-03-15 12:30:00', 5999.00, 5, 5, 1, 0);

-- Verify data
SELECT * FROM products;
SELECT * FROM flash_sales;
```

---

## Commit Strategy

- **All tasks**: `type(schema): create flash sale system tables` — seckill-database-schema.md

---

## Success Criteria Verification

```bash
# Verify schema syntax
mysql < seckill-database-schema.sql

# Verify all tables created
mysql -e "SHOW TABLES;"

# Verify indexes
mysql -e "SELECT TABLE_NAME, INDEX_NAME FROM information_schema.STATISTICS WHERE TABLE_SCHEMA = 'seckill_demo';"

# Verify triggers
mysql -e "SELECT TRIGGER_NAME, EVENT_MANIPULATION FROM information_schema.TRIGGERS WHERE TRIGGER_SCHEMA = 'seckill_demo';"
```
