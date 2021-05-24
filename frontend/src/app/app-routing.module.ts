import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { NotFoundComponent } from './not-found/not-found.component';
import { TableComponent } from './table/table.component';
import { UploadComponent } from './upload/upload.component';

import {  CanActivate } from '@angular/router';
import { AuthGuardGuard } from './guards/auth-guard.guard';
import { JwtModule, JwtModuleOptions } from '@auth0/angular-jwt';


const routes:  Routes = [
  { path:"", component:LoginComponent},
  { path: "table", component: TableComponent ,canActivate: [AuthGuardGuard] },
  { path: "upload", component: UploadComponent,canActivate: [AuthGuardGuard]},
  { path: "login", component:LoginComponent},
  { path: "**", component: NotFoundComponent}
  
]
export function tokenGetter() {
  return localStorage.getItem("access_token");
}
const JWT_Module_Options: JwtModuleOptions = {
  config: {
      tokenGetter: tokenGetter,
      allowedDomains: ["localhost:4200", "localhost:10000"]
  }
};


@NgModule({
  imports: [RouterModule.forRoot(routes),
            JwtModule.forRoot(JWT_Module_Options)],
  exports: [RouterModule],
  providers: [    
    AuthGuardGuard  
  ]
})
export class AppRoutingModule { }
