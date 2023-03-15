import { Component } from '@angular/core';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  providers: [HttpClient]
})

@Injectable({
  providedIn: 'root'
})

export class LoginComponent {
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

  async login() {
    const headers = new Headers( {
        'Content-Type': 'application/json'
    });
    const data = {FirstName: this.firstName, LastName: this.lastName, Email: this.email, Password: this.password}
    console.log(data);
    this.http.post('http://localhost:4200/api/login', data).toPromise();
    this.firstName = "";
    this.lastName = "";
    this.email = "";
    this.password = "";
    return;
  }
}

export interface LoginFields {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
}
