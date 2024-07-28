import { HttpContext, HttpHeaders, HttpParams } from '@angular/common/http';

export interface Options {
  headers?:
    | HttpHeaders
    | {
        [header: string]: string | string[];
      };
  observe?: 'body';
  context?: HttpContext;
  params?:
    | HttpParams
    | {
        [param: string]:
          | string
          | number
          | boolean
          | ReadonlyArray<string | number | boolean>;
      };
  reportProgress?: boolean;
  responseType?: 'json';
  withCredentials?: boolean;
  transferCache?:
    | {
        includeHeaders?: string[];
      }
    | boolean;
}

export interface NotesResponse {
  code: number,
  status: string,
  data: Note[]
}

export interface Note {
  id: number;
  text: string;
}

export interface PaginationParams {
  //[key: string]: string | number | boolean | ReadonlyArray<string | number | boolean>;
  page: number;
  perPage: number;
}
