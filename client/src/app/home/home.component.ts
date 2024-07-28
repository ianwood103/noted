import { Component } from '@angular/core';
import { NgIf, NgFor } from '@angular/common';
import { NoteService } from '../services/note.service';
import { Note, NotesResponse } from '../../types';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [NgIf, NgFor],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css',
})
export class HomeComponent {
  notes: Note[] | null = null;

  constructor(private noteService: NoteService) {}

  ngOnInit() {
    this.noteService
      .getNotes('http://localhost:8888/api/note')
      .subscribe((response: NotesResponse) => {
        if (response.code == 200) {
          this.notes = response.data;
        }
      });
  }
}
