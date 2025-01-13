USE inception;

-- ----------------------------
-- 用户基础信息表
-- ----------------------------
DROP TABLE IF EXISTS sys_member;
CREATE TABLE sys_member
(
    id        BIGINT AUTO_INCREMENT NOT NULL COMMENT '用户唯一标识符，主键',
    username  VARCHAR(64)           NOT NULL DEFAULT '' COMMENT '用户登录时使用的用户名',
    password  VARCHAR(100)          NOT NULL DEFAULT '' COMMENT '用户账户的密码，应存储加密后的值',
    nickname  VARCHAR(128)                   DEFAULT '' COMMENT '用户在系统中显示的名字，可选',
    avatar    VARCHAR(255)                   DEFAULT '' COMMENT '用户头像的URL或路径',
    quote     VARCHAR(255)                   DEFAULT '' COMMENT '用户设置的个人座右铭或签名',
    birthday  DATE                           DEFAULT '1970-01-01' COMMENT '用户的出生日期，用于个性化服务或统计',
    gender    INT                            DEFAULT 0 COMMENT '用户的性别，0表示保密，1表示男性，2表示女性',
    email     VARCHAR(128)                   DEFAULT '' COMMENT '用户的电子邮件地址，可用于找回密码或接收通知',
    telephone VARCHAR(20)                    DEFAULT '' COMMENT '用户的联系电话，可用于身份验证或联系用户',
    -- 固定字段
    status    INT                            DEFAULT 1 COMMENT '状态，0表示禁用，1表示正常启用',
    deleted   INT                            DEFAULT 0 COMMENT '逻辑删除标志，0表示未删除，1表示已删除，允许数据恢复',
    remark    VARCHAR(255)                   DEFAULT '' COMMENT '对记录的备注信息，如特殊说明等',
    create_at TIMESTAMP                      DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建的时间戳',
    create_by VARCHAR(64)                    DEFAULT '' COMMENT '创建该记录的用户标识符',
    update_at TIMESTAMP                      DEFAULT CURRENT_TIMESTAMP COMMENT '记录最后更新的时间戳',
    update_by VARCHAR(64)                    DEFAULT '' COMMENT '最后更新该记录的用户标识符',
    -- 添加索引
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT '用户表';