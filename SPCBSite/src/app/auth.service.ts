import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { JwtModule, JwtHelperService } from "@auth0/angular-jwt";
import { HttpClientModule } from "@angular/common/http";

@Injectable()

export class AuthService {

    helper = new JwtHelperService();
    
    constructor(){

    }

    loggedInMethod(){
        const token = localStorage.getItem('token');
        return !this.helper.isTokenExpired(token);
    }
}