interface SftpRequest {
  type: 'upload' | 'download' | 'list' | 'mkdir' | 'rmdir' | 'delete'; // 请求类型
  path: string
  file?: Buffer | Blob; // 文件内容（用于上传操作）
  recursive?: boolean; // 是否递归操作（用于目录）
}