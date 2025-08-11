export interface Session {
  // 会话ID
  id: string;
  // 会话名称
  labelName: string;
  // 会话主机
  host: string;
  // 会话端口
  port: number;
  // 会话用户名
  username: string;
  // 会话认证类型
  authType: 'password' | 'privateKey';
  // 会话密码
  password?: string;
  // 会话私钥
  privateKey?: string;
  // 会话私钥密码
  passphrase?: string;
  // 会话超时时间
  timeout: number;
  // 会话编码
  encoding: string;
  // 会话终端类型
  terminalType: string;
}
export interface SessionGroup {
  // 会话组ID
  id: string;
  // 会话组名称
  label: string;
  // 会话组会话列表
  sessionList: Session[];
}
