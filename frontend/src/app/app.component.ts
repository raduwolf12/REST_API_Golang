import { Component } from '@angular/core';
import { Validators } from '@angular/forms';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'lab2-project';
  isLogged:Boolean = false;
  verifyIsLogged(isLogged:Boolean)
  {
    this.isLogged = isLogged;
  }
}
