import { Component } from '@angular/core';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { AuthService } from '../auth.service';

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

  constructor(private http: HttpClient, public authService:AuthService) {
    this.firstName = "";
    this.lastName = "";
    this.email = "";
    this.password = "";
  }

  async login() {
    const headers = new Headers( {
        'Content-Type': 'application/json'
    });
    const data = {Email: this.email, Password: this.password}
    console.log(data);
    this.http.post('http://localhost:4200/api/login', data).toPromise();
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
