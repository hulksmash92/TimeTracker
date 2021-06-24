import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { HomeComponent } from './pages/home/home.component';
import { NotFoundComponent } from './pages/not-found/not-found.component';

const routes: Routes = [
  { 
    path: 'auth',
    loadChildren: () => import('./pages/auth/auth.module').then(m => m.AuthModule)
  },
  { path: '', component: HomeComponent },
  { path: '404', component: NotFoundComponent },
  { path: '**', redirectTo: '404' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
