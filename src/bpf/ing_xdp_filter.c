/*
*	Copyright (C) 2022  Kendall Tauser
*
*	This program is free software; you can redistribute it and/or modify
*	it under the terms of the GNU General Public License as published by
*	the Free Software Foundation; either version 2 of the License, or
*	(at your option) any later version.
*
*	This program is distributed in the hope that it will be useful,
*	but WITHOUT ANY WARRANTY; without even the implied warranty of
*	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*	GNU General Public License for more details.
*
*	You should have received a copy of the GNU General Public License along
*	with this program; if not, write to the Free Software Foundation, Inc.,
*	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

#include <xdp_filter.h>
#include <net/ethernet.h>
#include <bpf/bpf_helpers.h>
#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/in.h>
#include <linux/udp.h>

SEC("xdp1")
int ing_filter_xdp(struct xdp_md *xdp) {
    void *packet_end = (void *)(long)xdp->data_end;
    void *packet_begin = (void *)(long)xdp->data;

    int passl2 = filterl2(xdp, packet_end, packet_begin);
    
    switch(passl2) {
        case FORWARD_DROP: 
            return XDP_DROP;
        case FORWARD_RETURN:
            return XDP_TX;
        case FORWARD_UP:
            goto L3;
        case FORWARD_OUT:
            goto FORWARD;
    }
    
L3:
    int passl3 = filterl3(xdp, packet_end, packet_begin);
    
    switch(passl3) {
        case FORWARD_DROP:
            return XDP_DROP;
        case FORWARD_RETURN:
            return XDP_TX;
        case FORWARD_UP:
            goto L4;
        case FORWARD_OUT:
            goto FORWARD;
    }

L4:
    int passl4 = filterl4(xdp, packet_end, packet_begin);
    
    switch(passl4) {
        case FORWARD_DROP:
            return XDP_DROP;
        case FORWARD_RETURN:
            return XDP_TX;
        case FORWARD_UP:
            return XDP_PASS;
        case FORWARD_OUT:
            goto FORWARD;
    }   

FORWARD:
    return XDP_REDIRECT;
}

int filterl2(struct xdp_md *xdp, void *packet_end, void *packet_begin) {
    struct ethhdr *ehdr = packet_begin;

    return FORWARD_DROP;
}

int filterl3(struct xdp_md *xdp, void *packet_end, void *packet_begin) {
    struct ethhdr *ehdr = packet_begin;
    struct iphdr *ihdr = packet_begin + sizeof(ehdr);

    return FORWARD_DROP;
}

int filterl4(struct xdp_md *xdp, void *packet_end, void *packet_begin) {
    


    return FORWARD_DROP;
}


