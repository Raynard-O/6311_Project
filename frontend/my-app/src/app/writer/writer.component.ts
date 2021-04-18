import { Component, OnInit } from '@angular/core';
import { WriterService } from '../_services/writer.service';

@Component({
  selector: 'app-writer',
  templateUrl: './writer.component.html',
  styleUrls: ['./writer.component.css']
})
export class WriterComponent implements OnInit {

    shortLink: string = "";
    loading: boolean = false; // Flag variable
    file: File = null; // Variable to store file

  constructor(private writerService: WriterService) { }

  ngOnInit(): void {
  }
  onChange(event){
    this.file=event.target.files[0];
  }

  onUpload(){
    this.loading = !this.loading;
        console.log(this.file);
        this.writerService.upload(this.file).subscribe(
            (event: any) => {
                if (typeof (event) === 'object') {
  
                    // Short link via api response
                    this.shortLink = event.link;
  
                    this.loading = false; // Flag variable 
                }
            }
        );
  }

}
