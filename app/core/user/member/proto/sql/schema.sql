USE crucible_go;

-- ----------------------------
-- 角色表
-- ----------------------------
DROP TABLE IF EXISTS sys_role;
CREATE TABLE sys_role
(
    id        BIGINT AUTO_INCREMENT NOT NULL COMMENT '角色唯一标识符，主键',
    name      VARCHAR(50)  DEFAULT '' COMMENT '角色名称',
    code      VARCHAR(50)  DEFAULT '' COMMENT '角色编码',
    scope     VARCHAR(50)  DEFAULT '' COMMENT '角色范围',
    sort      INT          DEFAULT 0 COMMENT '排序',
    -- 固定字段
    status    INT          DEFAULT 1 COMMENT '状态，0表示禁用，1表示正常启用',
    deleted   INT          DEFAULT 0 COMMENT '逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复',
    remark    VARCHAR(255) DEFAULT '' COMMENT '对记录的备注信息，如特殊说明等',
    create_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间戳',
    create_by VARCHAR(64)  DEFAULT '' COMMENT '创建该记录的用户标识符',
    update_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新的时间戳',
    update_by VARCHAR(64)  DEFAULT '' COMMENT '最后更新该记录的用户标识符',
    -- 添加索引
    PRIMARY KEY (id),
    -- 添加唯一性约束
    CONSTRAINT unique_role_code UNIQUE (code)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT '角色表';


-- ----------------------------
-- 用户组表
-- ----------------------------
DROP TABLE IF EXISTS sys_group;
CREATE TABLE sys_group
(
    id        BIGINT AUTO_INCREMENT NOT NULL COMMENT '用户组唯一标识符，主键',
    name      VARCHAR(50)  DEFAULT '' COMMENT '用户组名称',
    code      VARCHAR(50)  DEFAULT '' COMMENT '用户组编码',
    parent_id BIGINT       DEFAULT 0 COMMENT '父级用户组ID',
    sort      INT          DEFAULT 0 COMMENT '排序',
    leader    BIGINT       DEFAULT 0 COMMENT '用户组负责人ID',
    -- 固定字段
    status    INT          DEFAULT 1 COMMENT '状态，0表示禁用，1表示正常启用',
    deleted   INT          DEFAULT 0 COMMENT '逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复',
    remark    VARCHAR(255) DEFAULT '' COMMENT '对记录的备注信息，如特殊说明等',
    create_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间戳',
    create_by VARCHAR(64)  DEFAULT '' COMMENT '创建该记录的用户标识符',
    update_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新的时间戳',
    update_by VARCHAR(64)  DEFAULT '' COMMENT '最后更新该记录的用户标识符',
    -- 添加索引
    PRIMARY KEY (id),
    -- 添加唯一性约束
    CONSTRAINT unique_group_code UNIQUE (code)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT '用户组表';


-- ----------------------------
-- 用户表
-- ----------------------------
DROP TABLE IF EXISTS sys_member;
CREATE TABLE sys_member
(
    id           BIGINT AUTO_INCREMENT NOT NULL COMMENT '用户唯一标识符，主键',
    identify     VARCHAR(64)           NOT NULL DEFAULT '' COMMENT '用户的唯一标识符',
    username     VARCHAR(64)           NOT NULL DEFAULT '' COMMENT '用户登录时使用的用户名',
    password     VARCHAR(100)          NOT NULL DEFAULT '' COMMENT '用户账户的密码，应存储加密后的值',
    nickname     VARCHAR(128)                   DEFAULT '' COMMENT '用户在系统中显示的名字，可选',
    avatar       VARCHAR(255)                   DEFAULT '' COMMENT '用户头像的URL或路径',
    quote        VARCHAR(255)                   DEFAULT '' COMMENT '用户设置的个人座右铭或签名',
    birthday     TIMESTAMP                      DEFAULT '1970-01-01 00:00:01' COMMENT '用户的出生日期，用于个性化服务或统计',
    gender       INT                            DEFAULT 0 COMMENT '用户的性别，0表示保密，1表示男性，2表示女性',
    email        VARCHAR(128)                   DEFAULT '' COMMENT '用户的电子邮件地址，可用于找回密码或接收通知',
    telephone    VARCHAR(20)                    DEFAULT '' COMMENT '用户的联系电话，可用于身份验证或联系用户',
    -- 记录字段
    login_at     TIMESTAMP                      DEFAULT CURRENT_TIMESTAMP COMMENT '用户最近一次登录的时间戳',
    login_ip     VARCHAR(64)                    DEFAULT '' COMMENT '用户最近一次登录的IP地址',
    login_region VARCHAR(64)                    DEFAULT '' COMMENT '用户最近一次登录的IP地址所属的地理位置',
    login_os     VARCHAR(64)                    DEFAULT '' COMMENT '用户最近一次登录的操作系统',
    -- 固定字段
    status       INT                            DEFAULT 1 COMMENT '状态，0表示禁用，1表示正常启用',
    deleted      INT                            DEFAULT 0 COMMENT '逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复',
    remark       VARCHAR(255)                   DEFAULT '' COMMENT '对记录的备注信息，如特殊说明等',
    create_at    TIMESTAMP                      DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间戳',
    create_by    VARCHAR(64)                    DEFAULT '' COMMENT '创建该记录的用户标识符',
    update_at    TIMESTAMP                      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新的时间戳',
    update_by    VARCHAR(64)                    DEFAULT '' COMMENT '最后更新该记录的用户标识符',
    -- 添加索引
    PRIMARY KEY (id),
    -- 添加唯一性约束
    CONSTRAINT unique_identify UNIQUE (identify),
    CONSTRAINT unique_username UNIQUE (username),
    CONSTRAINT unique_email UNIQUE (email),
    CONSTRAINT unique_telephone UNIQUE (telephone)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT '用户表';

-- ----------------------------
-- 角色用户关联表
-- ----------------------------
DROP TABLE IF EXISTS sys_member_role;
CREATE TABLE sys_member_role
(
    id        BIGINT AUTO_INCREMENT NOT NULL COMMENT '角色用户关联唯一标识符，主键',
    member_id BIGINT                NOT NULL COMMENT '用户唯一标识符',
    role_id   BIGINT                NOT NULL COMMENT '角色唯一标识符',
    -- 添加索引
    PRIMARY KEY (id),
    -- 添加唯一性约束
    CONSTRAINT unique_member_role UNIQUE (member_id, role_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT '角色用户关联表';


-- ----------------------------
-- 用户用户组关联表
-- ----------------------------
DROP TABLE IF EXISTS sys_member_group;
CREATE TABLE sys_member_group
(
    id        BIGINT AUTO_INCREMENT NOT NULL COMMENT '用户用户组关联唯一标识符，主键',
    member_id BIGINT                NOT NULL COMMENT '用户唯一标识符',
    group_id  BIGINT                NOT NULL COMMENT '用户组唯一标识符',
    -- 添加索引
    PRIMARY KEY (id),
    -- 添加唯一性约束
    CONSTRAINT unique_member_group UNIQUE (member_id, group_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT '用户用户组关联表';


