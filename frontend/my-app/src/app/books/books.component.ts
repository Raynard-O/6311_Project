import { Component, OnInit } from '@angular/core';
import { Book } from '../_models/book';
import { BookService } from '../_services/book.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-books',
  templateUrl: './books.component.html',
  styleUrls: ['./books.component.css'],
  providers: [BookService]
})

export class BooksComponent implements OnInit {
  title="Books For Everyone";
  books:Book[];
  selectedBook: Book;
  constructor(
    private bookService: BookService,
    private router: Router,
    ) { }
  getBooks(): void {
    this.bookService.getBooks().then(books => this.books = books);
  }
  ngOnInit(): void { 
    this.getBooks()     
  }

  onSelect(book: Book): void{
    this.selectedBook= book;
  }

  gotoDetail(): void {
    this.router.navigate(['/detail', this.selectedBook.id]);
  }  
}
