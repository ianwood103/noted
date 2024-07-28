import { Injectable } from '@angular/core';
import { ApiService } from './api.service';
import { Observable } from 'rxjs';
import { NotesResponse } from '../../types';

@Injectable({
  providedIn: 'root',
})
export class NoteService {
  constructor(private apiService: ApiService) {}

  getNotes = (url: string): Observable<NotesResponse> => {
    return this.apiService.get(url, { responseType: 'json' });
  };
}
