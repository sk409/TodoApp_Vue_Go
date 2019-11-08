import axios from "axios"
import {
    serverPortNumber
} from "../constants.js"

const origin = "http://localhost:" + serverPortNumber;

export default class Model {

    constructor(id) {
        this.id = id;
    }

    path(id) {
        const p = origin + "/" + this.basePath();
        return id ? p + "/" + id : p;
    }

    query() {
        let q = "";
        for (const key in this.params()) {
            q += `${key}=${this.params()[key]}&`;
        }
        return q;
    }

    store(handler) {
        const that = this;
        axios.post(this.path(), this.params()).then(response => {
            that.id = response.data.id;
            if (handler) {
                handler(response)
            }
        });
    }

    fetch(handler) {
        axios.get(this.path() + "?" + this.query()).then(response => {
            // TODO: データをセットする
            if (handler) {
                handler(response)
            }
        });
    }

    update() {
        axios.put(this.path(this.id), this.params());
    }

    delete() {
        axios.delete(this.path(this.id), this.params());
    }

    get id() {
        return this._id;
    }

    set id(value) {
        this._id = value;
    }

}