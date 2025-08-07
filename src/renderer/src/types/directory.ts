export interface FileItem{
  isDir : boolean;
  modTime: string;
  Mode: string;
  name: string;
  size: number;
}

export type FileList = FileItem[];

//用于文件操作的对象
export interface FileOperationObj {
  uniqId: string;
  fileList: FileList;
  currentPath: string;
  websocket: WebSocket;
}

export type FieleOperationObjList = FileOperationObj[] 