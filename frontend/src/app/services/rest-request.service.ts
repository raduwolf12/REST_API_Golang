import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { EventEmitter, Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { TokenServiceService } from './token-service.service';

@Injectable({
  providedIn: 'root'
})
export class RestRequestService {

  constructor(private http: HttpClient,
    private token: TokenServiceService) { }


  resolveItems(): Observable<HttpErrorResponse|any> {
    console.log('Request is sent!');
    return this.http.get("http://localhost:4200/").pipe(catchError(error=> this.handleError(error)));
  }

  url = "http://localhost:4200";
  post(user: { email: string, password: string }, url: string): Observable<any> {
    const event = new EventEmitter();

    setTimeout(() => {
      event.emit({ myObj: {} })
    })
    return event;
  }

  get(): Observable<HttpErrorResponse|any> {
    return this.http.get(this.url).pipe(catchError(error=> this.handleError(error)));;
  }
  get1(url1: string): Observable<HttpErrorResponse|any> {
    const headers = { headers: new HttpHeaders({ 'Content-Type': 'application/x-www-form-urlencoded', 'Access-Control-Allow-Origin': '*' }) };


    return this.http.get(url1, headers).pipe(catchError(error=> this.handleError(error)));
  }

  handleError(error: any) {
    if (error instanceof HttpErrorResponse && error.status === 401) {
      this.token.removeToken() 
    }
      return throwError(error)
  }


}
