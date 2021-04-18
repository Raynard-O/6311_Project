import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Location } from '@angular/common';
import {switchMap } from 'rxjs/operators';
import { BookService } from '../_services/book.service';
import { Book } from '../_models/book';

@Component({
  selector: 'app-book-details',
  templateUrl: './book-details.component.html',
  styleUrls: ['./book-details.component.css']
})
export class BookDetailsComponent implements OnInit {
  book: Book;

  constructor(
    private bookService: BookService,
    private route: ActivatedRoute,
    private location: Location
  ) { }
  ngOnInit(): void {
    this.route.paramMap.pipe(
    switchMap((params: Params) =>
      this.bookService.getBook(+params.get('id'))
      ))
    .subscribe(book => this.book  = book);
  }
  goBack(): void {
    this.location.back();
  }  
}