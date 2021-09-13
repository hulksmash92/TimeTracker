import { ComponentFixture } from '@angular/core/testing';
import { HarnessLoader, ComponentHarness, HarnessQuery } from '@angular/cdk/testing';
import { TestbedHarnessEnvironment } from '@angular/cdk/testing/testbed';
import { DebugElement, Type } from '@angular/core';
import { By } from '@angular/platform-browser';

/**
 * @description
 * Page class for testing a component's template with helper methods.
 * Subclass off of this to follow the Page object model whilst testing
 * 
 * @example // In your component spec file create a subclass like
 * class AppComponentTestingPage extends TestingPage {
 *     get button() { return this.query<HTMLButtonElement>('button.my-button'); }
 *     
 *     constructor(protected readonly fixture: ComponentFixture<AppComponent>) {
 *         super(fixture);
 *     }
 * }
 */
export abstract class TestingPage {
    /**
     * Material component harness loader
     */
    protected readonly matHarnessLoader: HarnessLoader;

    constructor(protected readonly fixture: ComponentFixture<any>) {
        if (!!fixture) {
            this.matHarnessLoader = TestbedHarnessEnvironment.loader(fixture);
        }
    }

    /**
     * Finds the first instance of an elemnt in the view with a specified css selector
     *
     * @param selector valid css selector to the intended element
     * 
     * @returns Element found with type `T`
     * 
     * @example const btn = this.query<HTMLButtonElement>('.btn');
     */
    protected query<T>(selector: string): T {
        return this.fixture.nativeElement.querySelector(selector);
    }

    /**
     * Finds all instances of an elemnt in the view with a specified css selector
     *
     * @param selector valid css selector to the intended element
     * 
     * @returns array of elements found with type `T` or an empty array if no items are found
     * 
     * @example const btns = this.queryAll<HTMLButtonElement>('.btn');
     */
    protected queryAll<T>(selector: string): T[] {
        return this.fixture.nativeElement.querySelectorAll(selector);
    }

    /**
     * Finds the first instance of an elemnt in the view with a specified css selector
     *
     * @param selector valid css selector to the intended element
     * 
     * @returns DebugElement of the selected element
     * 
     * @example const btn = this.queryByCss('.btn');
     */
    protected queryByCss(selector: string): DebugElement {
        return this.fixture.debugElement.query(By.css(selector));
    }

    /**
     * Finds all instances of an elemnt in the view with a specified css selector
     *
     * @param selector valid css selector to the intended element
     * 
     * @returns an array DebugElements of the selected elements or an empty array if no elements are found
     * 
     * @example const btns = this.queryAllByCss('.btn');
     */
    protected queryAllByCss(selector: string): DebugElement[] {
        return this.fixture.debugElement.queryAll(By.css(selector));
    }

    /**
     * Finds the first instance of a certain component in the view
     *
     * @param type directive type to find
     * 
     * @returns DebugElement of the selected component
     * 
     * @example const table = this.queryByDirective(MatTable);
     */
    protected queryByDirective<T>(type: Type<T>): DebugElement {
        return this.fixture.debugElement.query(By.directive(type));
    }

    /**
     * Finds the first instance of a certain component in the view
     *
     * @param type directive type to find
     * 
     * @returns array of DebugElements of the selected component or an empty list if no elements are found
     * 
     * @example const tables = this.queryAllByDirective(MatTable);
     */
    protected queryAllByDirective<T>(type: Type<T>): DebugElement[] {
        return this.fixture.debugElement.queryAll(By.directive(type));
    }

    /**
     * Wrapper function for the material harness loader's getHarness method
     * 
     * @param harnessQuery query for the harness to create
     */
    public getHarness<T extends ComponentHarness>(harnessQuery: HarnessQuery<T>): Promise<T> {
        return this.matHarnessLoader.getHarness(harnessQuery);
    }

    /**
     * Wrapper function for the material harness loader's getAllHarnesses method
     * 
     * @param harnessQuery query for the harness to create
     */
    public getAllHarnesses<T extends ComponentHarness>(harnessQuery: HarnessQuery<T>): Promise<T[]> {
        return this.matHarnessLoader.getAllHarnesses(harnessQuery);
    }
}
