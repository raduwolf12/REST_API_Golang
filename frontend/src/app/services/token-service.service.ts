import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class TokenServiceService {

  constructor() { }
  public setToken(token:string ){
    localStorage.setItem('token',token)
  } 
  public getToken():string{
  var tok =localStorage.getItem('token')
  
    if(tok!=null)
      return tok
   return ''
  }
  public removeToken(){
    localStorage.setItem('token','')
  }
}
