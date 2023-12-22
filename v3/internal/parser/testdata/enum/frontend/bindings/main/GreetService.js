// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

/**
 * @typedef {import('./models').Person} Person
 * @typedef {import('./models').Title} Title
 */

/**
 * Greet does XYZ 
 * @function Greet
 * @param name {string}
 * @param title {Title}
 * @returns {Promise<string>}
 **/
export function Greet(name, title) {
	return wails.CallByID(1411160069, ...Array.prototype.slice.call(arguments, 0));
}

/**
 * NewPerson creates a new person 
 * @function NewPerson
 * @param name {string}
 * @returns {Promise<Person>}
 **/
export function NewPerson(name) {
	return wails.CallByID(1661412647, ...Array.prototype.slice.call(arguments, 0));
}
