import { Component } from '@angular/core';
import { trigger, transition, style, animate } from '@angular/animations';

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css'],
  providers: [HttpClient]
})

@Injectable({
  providedIn: 'root'
})

export class SignupComponent {
  firstName: string;
  lastName: string;
  email: string;
  password: string;

  constructor(private http: HttpClient) {
    this.firstName = "";
    this.lastName = "";
    this.email = "";
    this.password = "";
  }

  async signup() {
    const headers = new Headers( {
        'Content-Type': 'application/json'
    });
    const data = {FirstName: this.firstName, LastName: this.lastName, Email: this.email, Password: this.password}
    console.log(data);
    this.http.post('http://localhost:4200/api/signup', data).toPromise();
    this.firstName = "";
    this.lastName = "";
    this.email = "";
    this.password = "";
    return;
  }
}

export interface SignUpFields {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
}
