import { Component } from '@angular/core';
import { NgIf, NgFor } from '@angular/common';
import { NoteService } from '../services/note.service';
import { Note, NotesResponse } from '../../types';
import { FormsModule } from '@angular/forms';

const generateNumberId = () => Math.floor(Math.random() * 1000000);

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [NgIf, NgFor, FormsModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css',
})
export class HomeComponent {
  notes: Note[] | null = null;

  constructor(private noteService: NoteService) {}

  ngOnInit() {
    this.noteService.getNotes('http://localhost:8888/api/note').subscribe({
      next: (response: NotesResponse) => (this.notes = response.data),
      error: (err) => console.error('Error fetching notes:', err),
    });
  }

  addNewNote(text: string) {
    const newNote: Note = { id: generateNumberId(), text: text };
    this.noteService
      .createNote('http://localhost:8888/api/note', newNote)
      .subscribe({
        next: () => {
          this.notes = this.notes ? [...this.notes, newNote] : [newNote];
          console.log('Note added successfully:', newNote);
        },
        error: (err) => console.error('Error adding note:', err),
      });
  }

  showModalFlag: boolean = false;
  newNoteText: string = '';

  showModal() {
    this.showModalFlag = true;
  }
}
