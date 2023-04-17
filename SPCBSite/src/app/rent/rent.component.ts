import { Component } from '@angular/core';

@Component({
  selector: 'app-rent',
  templateUrl: './rent.component.html',
  styleUrls: ['./rent.component.css']
})


export class RentComponent {
  firstName: string;
  lastName: string;
  email: string;
  password: string;

  constructor() {
    this.firstName = "";
    this.lastName = "";
    this.email = "";
    this.password = "";
  }
}
