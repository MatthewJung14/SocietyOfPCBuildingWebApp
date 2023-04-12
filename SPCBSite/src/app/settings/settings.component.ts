import { Component } from '@angular/core';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css'],
  providers: [AuthService],
})

@Injectable({
  providedIn: 'root'
})

export class SettingsComponent {
  edit = 'default';
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

  async updateAccount(change: string) {
    const headers = new Headers( {
        'Content-Type': 'application/json'
    });
    if (change == "name"){
      const data = {FirstName: this.firstName, LastName: this.lastName}
      console.log(data);
      this.http.put('http://localhost:4200/api/update-account', data).toPromise();
      this.firstName = "";
      this.lastName = "";
    } else if (change == "password") {
      const data = {Password: this.password}
      console.log(data);
      this.http.put('http://localhost:4200/api/update-account', data).toPromise();
      this.password = "";
    } else {
      const data = {Email: this.email}
      console.log(data);
      this.http.put('http://localhost:4200/api/update-account', data).toPromise();
      this.email = "";
    }

    return;
  }

  logout() {
    localStorage.removeItem('token');
    this.router.navigate(['home']);
  }
}

export interface settingsFields {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
}