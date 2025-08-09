export const MessageType = {
  FITNESS: 'fitness',
  READ_SESSIONS: 'readSessions',
}
export interface fitness {
  row: number,
  col: number,
}

export interface request {
  reqType: string,
  content?: fitness,
}
