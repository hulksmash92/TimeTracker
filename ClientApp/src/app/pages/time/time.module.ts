import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { TimeComponent } from './time.component';
import { FormComponent } from './form/form.component';
import { TableComponent } from './table/table.component';
import { TimeRoutingModule } from './time-routing.module';
import { MaterialModule } from 'src/app/modules/material/material.module';

@NgModule({
  declarations: [
    TimeComponent,
    FormComponent,
    TableComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    MaterialModule,
    TimeRoutingModule
  ]
})
export class TimeModule { }
