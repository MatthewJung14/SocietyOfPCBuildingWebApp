import { Component, Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  providers: [AuthService, HttpClient],
})
export class LoginComponent {
  firstName: string;
  lastName: string;
  email: string;
  password: string;

  constructor(private http: HttpClient, public authService: AuthService, private router: Router) {
    this.firstName = "";
    this.lastName = "";
    this.email = "";
    this.password = "";
  }

  async login() {
    console.log('clicked')
    const headers = new Headers({
      'Content-Type': 'application/json'
    });
    const data = { Email: this.email, Password: this.password };
    console.log('data' + data);
    const response = this.http.post('http://localhost:4200/api/login', data).toPromise();
    // const jsonResponse = JSON.parse(response); // parse the JSON response string
    // const token = jsonResponse['token'];
    this.email = "";
    this.password = "";
    console.log(response)
    const string = JSON.stringify(response);
    //const responseBody = JSON.parse(stringifyWithZone(response));
    //console.log('responseBody', responseBody);
    console.log('stringified'+string);
    localStorage.setItem('token', JSON.stringify(data));
    this.router.navigate(['home']);
    return;
  }

  async forgot() {
    const headers = new Headers({
      'Content-Type': 'application/json'
    });
    const data = { Email: this.email, Password: this.password };
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
