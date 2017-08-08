
import fetch from 'isomorphic-fetch'

export function PayAction(orderId) {
    let request_data = {
        OrderID: orderId,
        Status: "orderPay"
    }

    fetch("/api/order_update", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json, text/plain, */*'
        },
        body: JSON.stringify(request_data),
        credentials: 'same-origin'
    }).then(data => data.json()).then(data=>{
        return data
    })
}

export function CompleteAction(orderId) {
    let request_data = {
        OrderID: orderId,
        Status: "orderComplete"
    }
    fetch("/api/order_update", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json, text/plain, */*'
        },
        body: JSON.stringify(request_data),
        credentials: 'same-origin'
    }).then(data => data.json()).then(data=>{
        return data
    })
}

export function StatusAction(orderId, status) {
    let result = null;
    if(status == 0)  //未支付
    {
        result = PayAction(orderId)
    } else if (status == 1) //已经购买
    {
        result = CompleteAction(orderId)
    }
    if(result.ok == true) {
        location.reload(true)
    }

}

