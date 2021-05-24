import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { RestRequestService } from './rest-request.service';
import {tap} from 'rxjs/operators'
import { User } from '../models/User';

@Injectable({
  providedIn: 'root'
})
export class UserserviceService {

  constructor(public restRequestService: RestRequestService) { }

  userData!: User;
   
  login2(email: string, password: string) {
    let postData = {email : email ,password :password};
    console.log('Request is sent!');
    // Using the POST method
    const headers =  {headers: new  HttpHeaders({ 'Content-Type': 'application/x-www-form-urlencoded'})};
    return this.restRequestService.get1( "http://localhost:10000/login/"+email+"/"+password).pipe(tap(
              data =>{ this.userData=data
              console.log(data)}
            ));
  }
  
  

  

}
