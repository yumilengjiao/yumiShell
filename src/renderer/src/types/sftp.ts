export interface FileObj {
  name: string;         // 文件名
  content: string;      // Base64编码的文件内容
}

export interface SftpRequest {
  requestType: 'get' | 'upload' | 'download' | 'mkdir' | 'rmdir' | 'delete' | "create"; // 请求类型
  path: string
  file?: FileObj;
  recursive?: boolean; // 是否递归操作（用于目录）
}