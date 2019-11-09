import axios from "axios"
import {
    serverPortNumber
} from "../constants.js"

const origin = "http://localhost:" + serverPortNumber;

export default class Model {

    static fetch(basePath, params, handler) {
        var url = origin + "/" + basePath;
        if (params) {
            url += "?" + Model.query(params);
        }
        axios.get(url).then(response => {
            // TODO: データをセットする
            if (handler) {
                handler(response);
            }
        });
    }

    static query(params) {
        let q = "";
        for (const key in params) {
            q += `${key}=${params[key]}&`;
        }
        return q;
    }

    constructor(id) {
        this.id = id;
    }

    path(id) {
        const p = origin + "/" + this.basePath();
        return id ? p + "/" + id : p;
    }

    store(handler) {
        const that = this;
        axios.post(this.path(), this.params()).then(response => {
            that.id = response.data.id;
            if (handler) {
                handler(response);
            }
        });
    }

    update(handler) {
        axios.put(this.path(this.id), this.params()).then(response => {
            if (handler) {
                handler(response);
            }
        });
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