import { Component } from '@angular/core';
import { trigger, transition, style, animate } from '@angular/animations';

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css'],
  providers: [AuthService, HttpClient]
})

@Injectable({
  providedIn: 'root'
})

export class SignupComponent {
  firstName: string;
  lastName: string;
  email: string;
  password: string;

  constructor(private http: HttpClient, private router: Router, public authService: AuthService) {
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
    const response = this.http.post('http://localhost:4200/api/signup', data).toPromise();
    console.log(response)
    console.log('stringified'+JSON.stringify(data));
    this.firstName = "";
    this.lastName = "";
    this.email = "";
    this.password = "";
    this.router.navigate(['login']);
    return;
  }
}

export interface SignUpFields {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
}
