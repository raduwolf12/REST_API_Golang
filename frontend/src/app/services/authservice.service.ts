import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { TokenServiceService } from './token-service.service';
@Injectable({
  providedIn: 'root'
})
export class AuthService {
  
  constructor(public jwtHelper: JwtHelperService, public tokenService: TokenServiceService) {}
  // ...
  public isAuthenticated(): boolean {
    const token = this.tokenService.getToken();
    // const token = localStorage.getItem('token');
    // Check whether the token is expired and return
    // true or false
  if(token=='')
  {
    return false;
  }

    return !this.jwtHelper.isTokenExpired(token?.toString());
  }
}