import { Component } from '@angular/core';
import { NoteService } from '../services/note.service';
import { NotesResponse } from '../../types';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css',
})
export class HomeComponent {
  constructor(private noteService: NoteService) {}

  ngOnInit() {
    this.noteService
      .getNotes('http://localhost:8888/api/note')
      .subscribe((notes: NotesResponse) => {
        console.log(notes);
      });
  }
}
