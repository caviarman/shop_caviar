import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { User } from '../models/user.model'; 


import { Observable, of } from 'rxjs';
import { catchError } from 'rxjs/operators';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type':  'application/json',
    Authorization: 'my-auth-token'
  })
};

@Injectable()
export class UserService {
  usersUrl = `${environment.host}/api/user`;  

  constructor(private http: HttpClient) {}

  getUsers(): Observable<any> {
    return this.http.get(this.usersUrl)
      .pipe(
        catchError(err => of('error', err))
      );
  }

  crateUser(user: User): Observable<any> {
    return this.http.post(this.usersUrl, user)
      .pipe(
        catchError(err => of('error', err))
      );
  }
}