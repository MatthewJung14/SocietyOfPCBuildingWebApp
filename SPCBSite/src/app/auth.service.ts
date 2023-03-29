import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { JwtModule, JwtHelperService } from "@auth0/angular-jwt";
import { HttpClientModule } from "@angular/common/http";
import * as jwt from 'jsonwebtoken';

@Injectable()

export class AuthService {

    helper = new JwtHelperService();
    
    constructor(private http: HttpClient){}

    loggedInMethod(){
        const token = localStorage.getItem('token');
        if (token === null){
            console.log("false")
            return false
        } else {
            console.log("true")
            return true
        }
        //const decodedToken = jwt.decode(token);
        //console.log(!this.helper.isTokenExpired(decodedToken))
        //return !this.helper.isTokenExpired(decodedToken);
    }
}

        // if (localStorage.getItem('token'))
        // return this.isTokenExpired(token);
    //     if (token == null){
    //         console.log("false")
    //         return false
    //     } else {
    //         console.log("true")
    //         return true
    //     }