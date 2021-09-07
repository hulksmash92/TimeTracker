import { Component } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'user-delete-confirm',
  template: `
    <h1 mat-dialog-title>Are you sure?</h1>
    <mat-dialog-content>
        <p>Are you sure you want to delete your account?</p> 
        <p>This will only delete your data from this application</p>
    </mat-dialog-content>
    <mat-dialog-actions align="end">
        <button mat-button 
                color="warn" 
                class="mr-2" [
                (click)="closeDialog(true)">
            Yes
        </button>
        <button mat-button 
                color="accent" 
                class="ml-2"
                (click)="closeDialog(false)">
            Cancel
        </button>
    </mat-dialog-actions>
  `
})
export class UserDeleteConfirmComponent {

  constructor(public matDialogRef: MatDialogRef<UserDeleteConfirmComponent>) { }

  closeDialog(confirm: boolean): void {
    this.matDialogRef.close(confirm);
  }
}
