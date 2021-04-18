import { Injectable } from '@angular/core';
import { Book } from '../_models/book';

const Books: Book[] = [
  { id: 1, name: 'Harry Potter', author: 'J.K.Rowling', Chapter: '221', Content: 'Mr. and Mrs. Dursley, of number four, Privet Drive, were proud to say that they were perfectly normal, thank you very much. They were the last people you’d expect to be involved in anything strange or mysterious, because they just didn’t hold with such nonsense. Mr. Dursley was the director of a firm called Grunnings, which made drills. He was a big, beefy man with hardly any neck, although he did have a very large mustache. Mrs. Dursley was thin and blonde and had nearly twice the usual amount of neck, which came in very useful as she spent so much of her time craning over garden fences, spying on the neighbors. The Dursleys had a small son called Dudley and in their opinion there was no finer boy anywhere. The Dursleys had everything they wanted, but they also had a secret, and their greatest fear was that somebody would discover it. They didn’t think they could bear it if anyone found out about the Potters.' },
  { id: 2, name: 'Twilight', author: '', Chapter: '',  Content: '' },
  { id: 3, name: 'New Moon', author: '', Chapter: '',  Content: '' },
  { id: 4, name: 'Eclipse', author: '', Chapter: '',  Content: '' },
  { id: 5, name: 'Breaking Dawn', author: '', Chapter: '',  Content: '' },
  { id: 6, name: 'Twenties Girl', author: '', Chapter: '',  Content: '' },
  { id: 7, name: 'Host', author: '', Chapter: '',  Content: '' },
  { id: 8, name: 'Remember Me', author: '', Chapter: '',  Content: '' },
  { id: 9, name: 'Finding Audrey', author: '', Chapter: '',  Content: '' },
  { id: 10, name: 'Angels and Demons', author: '', Chapter: '',  Content: '' }
];

@Injectable()
export class BookService {
  constructor() {
  }
  getBooks(): Promise<Book[]>{
    return Promise.resolve(Books);
  }
  getBook(id: number): Promise<Book> {
    return this.getBooks()
    .then(books => books.find(book => book.id === id));
  }
}

