import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MaterialModule } from './modules/material/material.module';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { UserMenuComponent } from './components/user-menu/user-menu.component';
import { AvatarModule } from './components/avatar/avatar.module';
import { NotFoundComponent } from './pages/not-found/not-found.component';
import { HomeComponent } from './pages/home/home.component';

@NgModule({
  declarations: [
    AppComponent,
    UserMenuComponent,
    NotFoundComponent,
    HomeComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    BrowserAnimationsModule,
    FormsModule,
    ReactiveFormsModule,
    AppRoutingModule,
    MaterialModule,
    AvatarModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
