CREATE DATABASE IF NOT EXISTS crucible_go;

USE crucible_go;

-- ----------------------------
-- 用户表
-- ----------------------------
DROP TABLE IF EXISTS sys_user;
CREATE TABLE sys_user
(
    id           BIGINT AUTO_INCREMENT NOT NULL COMMENT '用户唯一标识符，主键',
    identify     VARCHAR(64)           NOT NULL DEFAULT '' COMMENT '用户的唯一标识符',
    username     VARCHAR(64)           NOT NULL DEFAULT '' COMMENT '用户登录时使用的用户名',
    password     VARCHAR(100)          NOT NULL DEFAULT '' COMMENT '用户账户的密码，应存储加密后的值',
    nickname     VARCHAR(128)          NULL     DEFAULT NULL COMMENT '用户在系统中显示的名字，可选',
    avatar       VARCHAR(255)          NULL     DEFAULT NULL COMMENT '用户头像的URL或路径',
    quote        VARCHAR(255)          NULL     DEFAULT NULL COMMENT '用户设置的个人座右铭或签名',
    birthday     TIMESTAMP             NULL     DEFAULT NULL COMMENT '用户的出生日期，用于个性化服务或统计',
    gender       INT                   NULL     DEFAULT 0 COMMENT '用户的性别，0表示保密，1表示男性，2表示女性',
    email        VARCHAR(128)          NULL     DEFAULT NULL COMMENT '用户的电子邮件地址，可用于找回密码或接收通知',
    telephone    VARCHAR(20)           NULL     DEFAULT NULL COMMENT '用户的联系电话，可用于身份验证或联系用户',
    -- 记录字段
    login_at     TIMESTAMP             NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '用户最近一次登录的时间戳',
    login_ip     VARCHAR(64)           NULL     DEFAULT NULL COMMENT '用户最近一次登录的IP地址',
    login_region VARCHAR(64)           NULL     DEFAULT NULL COMMENT '用户最近一次登录的IP地址所属的地理位置',
    login_os     VARCHAR(64)           NULL     DEFAULT NULL COMMENT '用户最近一次登录的操作系统',
    -- 固定字段
    status       INT                   NULL     DEFAULT 1 COMMENT '状态，0表示禁用，1表示正常启用',
    deleted      INT                   NULL     DEFAULT 0 COMMENT '逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复',
    remark       VARCHAR(255)          NULL     DEFAULT NULL COMMENT '对记录的备注信息，如特殊说明等',
    create_at    TIMESTAMP             NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间戳',
    create_by    VARCHAR(64)           NULL     DEFAULT NULL COMMENT '创建该记录的用户标识符',
    update_at    TIMESTAMP             NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新的时间戳',
    update_by    VARCHAR(64)           NULL     DEFAULT NULL COMMENT '最后更新该记录的用户标识符',
    -- 添加索引
    PRIMARY KEY (id),
    -- 添加唯一性约束
    CONSTRAINT unique_identify UNIQUE (identify),
    CONSTRAINT unique_username UNIQUE (username),
    CONSTRAINT unique_email UNIQUE (email),
    CONSTRAINT unique_telephone UNIQUE (telephone)
) COMMENT '用户表';


-- ----------------------------
-- 用户组表
-- ----------------------------
DROP TABLE IF EXISTS sys_group;
CREATE TABLE sys_group
(
    id        BIGINT AUTO_INCREMENT NOT NULL COMMENT '用户组唯一标识符，主键',
    name      VARCHAR(50)           NULL DEFAULT NULL COMMENT '用户组名称',
    code      VARCHAR(50)           NULL DEFAULT NULL COMMENT '用户组编码',
    parent_id BIGINT                NULL DEFAULT 0 COMMENT '父级用户组ID',
    sort      INT                   NULL DEFAULT 0 COMMENT '排序',
    leader    BIGINT                NULL DEFAULT 0 COMMENT '用户组负责人ID',
    -- 固定字段
    status    INT                   NULL DEFAULT 1 COMMENT '状态，0表示禁用，1表示正常启用',
    deleted   INT                   NULL DEFAULT 0 COMMENT '逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复',
    remark    VARCHAR(255)          NULL DEFAULT NULL COMMENT '对记录的备注信息，如特殊说明等',
    create_at TIMESTAMP             NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间戳',
    create_by VARCHAR(64)           NULL DEFAULT NULL COMMENT '创建该记录的用户标识符',
    update_at TIMESTAMP             NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新的时间戳',
    update_by VARCHAR(64)           NULL DEFAULT NULL COMMENT '最后更新该记录的用户标识符',
    -- 添加索引
    PRIMARY KEY (id),
    -- 添加唯一性约束
    CONSTRAINT unique_group_code UNIQUE (code)
) COMMENT '用户组表';

-- ----------------------------
-- 用户用户组关联表
-- ----------------------------
DROP TABLE IF EXISTS sys_user_group;
CREATE TABLE sys_user_group
(
    id       BIGINT AUTO_INCREMENT NOT NULL COMMENT '用户用户组关联唯一标识符，主键',
    user_id  BIGINT                NOT NULL COMMENT '用户唯一标识符',
    group_id BIGINT                NOT NULL COMMENT '用户组唯一标识符',
    -- 添加索引
    PRIMARY KEY (id),
    -- 添加唯一性约束
    CONSTRAINT unique_user_group UNIQUE (user_id, group_id)
) COMMENT '用户用户组关联表';

-- ----------------------------
-- 角色表
-- ----------------------------
DROP TABLE IF EXISTS sys_role;
CREATE TABLE sys_role
(
    id        BIGINT AUTO_INCREMENT NOT NULL COMMENT '角色唯一标识符，主键',
    name      VARCHAR(50)           NULL DEFAULT NULL COMMENT '角色名称',
    code      VARCHAR(50)           NULL DEFAULT NULL COMMENT '角色编码',
    scope     VARCHAR(50)           NULL DEFAULT NULL COMMENT '角色范围',
    sort      INT                   NULL DEFAULT 0 COMMENT '排序',
    -- 固定字段
    status    INT                   NULL DEFAULT 1 COMMENT '状态，0表示禁用，1表示正常启用',
    deleted   INT                   NULL DEFAULT 0 COMMENT '逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复',
    remark    VARCHAR(255)          NULL DEFAULT NULL COMMENT '对记录的备注信息，如特殊说明等',
    create_at TIMESTAMP             NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间戳',
    create_by VARCHAR(64)           NULL DEFAULT NULL COMMENT '创建该记录的用户标识符',
    update_at TIMESTAMP             NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新的时间戳',
    update_by VARCHAR(64)           NULL DEFAULT NULL COMMENT '最后更新该记录的用户标识符',
    -- 添加索引
    PRIMARY KEY (id),
    -- 添加唯一性约束
    CONSTRAINT unique_role_code UNIQUE (code)
) COMMENT '角色表';

-- ----------------------------
-- 角色用户关联表
-- ----------------------------
DROP TABLE IF EXISTS sys_user_role;
CREATE TABLE sys_user_role
(
    id      BIGINT AUTO_INCREMENT NOT NULL COMMENT '角色用户关联唯一标识符，主键',
    user_id BIGINT                NOT NULL COMMENT '用户唯一标识符',
    role_id BIGINT                NOT NULL COMMENT '角色唯一标识符',
    -- 添加索引
    PRIMARY KEY (id),
    -- 添加唯一性约束
    CONSTRAINT unique_user_role UNIQUE (user_id, role_id)
) COMMENT '角色用户关联表';

-- ----------------------------
-- 字典类型表
-- ----------------------------
DROP TABLE IF EXISTS sys_dict_type;
CREATE TABLE sys_dict_type
(
    id        BIGINT AUTO_INCREMENT NOT NULL COMMENT '字典类型唯一标识符，主键',
    name      VARCHAR(50)           NULL DEFAULT NULL COMMENT '字典类型名称',
    code      VARCHAR(50)           NULL DEFAULT NULL COMMENT '字典类型编码',
    -- 固定字段
    status    INT                   NULL DEFAULT 1 COMMENT '状态，0表示禁用，1表示正常启用',
    deleted   INT                   NULL DEFAULT 0 COMMENT '逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复',
    remark    VARCHAR(255)          NULL DEFAULT NULL COMMENT '对记录的备注信息，如特殊说明等',
    create_at TIMESTAMP             NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间戳',
    create_by VARCHAR(64)           NULL DEFAULT NULL COMMENT '创建该记录的用户标识符',
    update_at TIMESTAMP             NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新的时间戳',
    update_by VARCHAR(64)           NULL DEFAULT NULL COMMENT '最后更新该记录的用户标识符',
    -- 添加索引
    PRIMARY KEY (id)
) COMMENT '字典类型表';

-- ----------------------------
-- 字典数据表
-- ----------------------------
DROP TABLE IF EXISTS sys_dict_data;
CREATE TABLE sys_dict_data
(
    id        BIGINT AUTO_INCREMENT NOT NULL COMMENT '字典数据唯一标识符，主键',
    name      VARCHAR(50)           NULL DEFAULT NULL COMMENT '字典数据名称',
    code      VARCHAR(50)           NULL DEFAULT NULL COMMENT '字典数据编码',
    type_id   BIGINT                NOT NULL COMMENT '字典类型唯一标识符',
    sort      INT                   NULL DEFAULT 0 COMMENT '排序',
    -- 固定字段
    status    INT                   NULL DEFAULT 1 COMMENT '状态，0表示禁用，1表示正常启用',
    deleted   INT                   NULL DEFAULT 0 COMMENT '逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复',
    remark    VARCHAR(255)          NULL DEFAULT NULL COMMENT '对记录的备注信息，如特殊说明等',
    create_at TIMESTAMP             NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间戳',
    create_by VARCHAR(64)           NULL DEFAULT NULL COMMENT '创建该记录的用户标识符',
    update_at TIMESTAMP             NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新的时间戳',
    update_by VARCHAR(64)           NULL DEFAULT NULL COMMENT '最后更新该记录的用户标识符',
    -- 添加索引
    PRIMARY KEY (id)
) COMMENT '字典数据表';

-- ----------------------------
-- 系统日志表
-- ----------------------------
DROP TABLE IF EXISTS sys_log;
CREATE TABLE sys_log
(
    id        BIGINT AUTO_INCREMENT NOT NULL COMMENT '日志唯一标识符，主键',

    -- 固定字段
    status    INT                   NULL DEFAULT 1 COMMENT '状态，0表示禁用，1表示正常启用',
    deleted   INT                   NULL DEFAULT 0 COMMENT '逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复',
    remark    VARCHAR(255)          NULL DEFAULT NULL COMMENT '对记录的备注信息，如特殊说明等',
    create_at TIMESTAMP             NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间戳',
    create_by VARCHAR(64)           NULL DEFAULT NULL COMMENT '创建该记录的用户标识符',
    update_at TIMESTAMP             NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新的时间戳',
    update_by VARCHAR(64)           NULL DEFAULT NULL COMMENT '最后更新该记录的用户标识符',
    -- 添加索引
    PRIMARY KEY (id)
) COMMENT '系统日志表';