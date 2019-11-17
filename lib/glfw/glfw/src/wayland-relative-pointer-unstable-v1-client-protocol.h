/* Generated by wayland-scanner 1.14.0 */

#ifndef RELATIVE_POINTER_UNSTABLE_V1_CLIENT_PROTOCOL_H
#define RELATIVE_POINTER_UNSTABLE_V1_CLIENT_PROTOCOL_H

#include <stdint.h>
#include <stddef.h>
#include "wayland-client.h"

#ifdef  __cplusplus
extern "C" {
#endif

/**
 * @page page_relative_pointer_unstable_v1 The relative_pointer_unstable_v1 protocol
 * protocol for relative pointer motion events
 *
 * @section page_desc_relative_pointer_unstable_v1 Description
 *
 * This protocol specifies a set of interfaces used for making clients able to
 * receive relative pointer events not obstructed by barriers (such as the
 * monitor edge or other pointer barriers).
 *
 * To start receiving relative pointer events, a client must first bind the
 * global interface "wp_relative_pointer_manager" which, if a compositor
 * supports relative pointer motion events, is exposed by the registry. After
 * having created the relative pointer manager proxy object, the client uses
 * it to create the actual relative pointer object using the
 * "get_relative_pointer" request given a wl_pointer. The relative pointer
 * motion events will then, when applicable, be transmitted via the proxy of
 * the newly created relative pointer object. See the documentation of the
 * relative pointer interface for more details.
 *
 * Warning! The protocol described in this file is experimental and backward
 * incompatible changes may be made. Backward compatible changes may be added
 * together with the corresponding interface version bump. Backward
 * incompatible changes are done by bumping the version number in the protocol
 * and interface names and resetting the interface version. Once the protocol
 * is to be declared stable, the 'z' prefix and the version number in the
 * protocol and interface names are removed and the interface version number is
 * reset.
 *
 * @section page_ifaces_relative_pointer_unstable_v1 Interfaces
 * - @subpage page_iface_zwp_relative_pointer_manager_v1 - get relative pointer objects
 * - @subpage page_iface_zwp_relative_pointer_v1 - relative pointer object
 * @section page_copyright_relative_pointer_unstable_v1 Copyright
 * <pre>
 *
 * Copyright © 2014      Jonas Ådahl
 * Copyright © 2015      Red Hat Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a
 * copy of this software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice (including the next
 * paragraph) shall be included in all copies or substantial portions of the
 * Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
 * THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
 * DEALINGS IN THE SOFTWARE.
 * </pre>
 */
struct wl_pointer;
struct zwp_relative_pointer_manager_v1;
struct zwp_relative_pointer_v1;

/**
 * @page page_iface_zwp_relative_pointer_manager_v1 zwp_relative_pointer_manager_v1
 * @section page_iface_zwp_relative_pointer_manager_v1_desc Description
 *
 * A global interface used for getting the relative pointer object for a
 * given pointer.
 * @section page_iface_zwp_relative_pointer_manager_v1_api API
 * See @ref iface_zwp_relative_pointer_manager_v1.
 */
/**
 * @defgroup iface_zwp_relative_pointer_manager_v1 The zwp_relative_pointer_manager_v1 interface
 *
 * A global interface used for getting the relative pointer object for a
 * given pointer.
 */
extern const struct wl_interface zwp_relative_pointer_manager_v1_interface;
/**
 * @page page_iface_zwp_relative_pointer_v1 zwp_relative_pointer_v1
 * @section page_iface_zwp_relative_pointer_v1_desc Description
 *
 * A wp_relative_pointer object is an extension to the wl_pointer interface
 * used for emitting relative pointer events. It shares the same focus as
 * wl_pointer objects of the same seat and will only emit events when it has
 * focus.
 * @section page_iface_zwp_relative_pointer_v1_api API
 * See @ref iface_zwp_relative_pointer_v1.
 */
/**
 * @defgroup iface_zwp_relative_pointer_v1 The zwp_relative_pointer_v1 interface
 *
 * A wp_relative_pointer object is an extension to the wl_pointer interface
 * used for emitting relative pointer events. It shares the same focus as
 * wl_pointer objects of the same seat and will only emit events when it has
 * focus.
 */
extern const struct wl_interface zwp_relative_pointer_v1_interface;

#define ZWP_RELATIVE_POINTER_MANAGER_V1_DESTROY 0
#define ZWP_RELATIVE_POINTER_MANAGER_V1_GET_RELATIVE_POINTER 1


/**
 * @ingroup iface_zwp_relative_pointer_manager_v1
 */
#define ZWP_RELATIVE_POINTER_MANAGER_V1_DESTROY_SINCE_VERSION 1
/**
 * @ingroup iface_zwp_relative_pointer_manager_v1
 */
#define ZWP_RELATIVE_POINTER_MANAGER_V1_GET_RELATIVE_POINTER_SINCE_VERSION 1

/** @ingroup iface_zwp_relative_pointer_manager_v1 */
static inline void
zwp_relative_pointer_manager_v1_set_user_data(struct zwp_relative_pointer_manager_v1 *zwp_relative_pointer_manager_v1, void *user_data)
{
	wl_proxy_set_user_data((struct wl_proxy *) zwp_relative_pointer_manager_v1, user_data);
}

/** @ingroup iface_zwp_relative_pointer_manager_v1 */
static inline void *
zwp_relative_pointer_manager_v1_get_user_data(struct zwp_relative_pointer_manager_v1 *zwp_relative_pointer_manager_v1)
{
	return wl_proxy_get_user_data((struct wl_proxy *) zwp_relative_pointer_manager_v1);
}

static inline uint32_t
zwp_relative_pointer_manager_v1_get_version(struct zwp_relative_pointer_manager_v1 *zwp_relative_pointer_manager_v1)
{
	return wl_proxy_get_version((struct wl_proxy *) zwp_relative_pointer_manager_v1);
}

/**
 * @ingroup iface_zwp_relative_pointer_manager_v1
 *
 * Used by the client to notify the server that it will no longer use this
 * relative pointer manager object.
 */
static inline void
zwp_relative_pointer_manager_v1_destroy(struct zwp_relative_pointer_manager_v1 *zwp_relative_pointer_manager_v1)
{
	wl_proxy_marshal((struct wl_proxy *) zwp_relative_pointer_manager_v1,
			 ZWP_RELATIVE_POINTER_MANAGER_V1_DESTROY);

	wl_proxy_destroy((struct wl_proxy *) zwp_relative_pointer_manager_v1);
}

/**
 * @ingroup iface_zwp_relative_pointer_manager_v1
 *
 * Create a relative pointer interface given a wl_pointer object. See the
 * wp_relative_pointer interface for more details.
 */
static inline struct zwp_relative_pointer_v1 *
zwp_relative_pointer_manager_v1_get_relative_pointer(struct zwp_relative_pointer_manager_v1 *zwp_relative_pointer_manager_v1, struct wl_pointer *pointer)
{
	struct wl_proxy *id;

	id = wl_proxy_marshal_constructor((struct wl_proxy *) zwp_relative_pointer_manager_v1,
			 ZWP_RELATIVE_POINTER_MANAGER_V1_GET_RELATIVE_POINTER, &zwp_relative_pointer_v1_interface, NULL, pointer);

	return (struct zwp_relative_pointer_v1 *) id;
}

/**
 * @ingroup iface_zwp_relative_pointer_v1
 * @struct zwp_relative_pointer_v1_listener
 */
struct zwp_relative_pointer_v1_listener {
	/**
	 * relative pointer motion
	 *
	 * Relative x/y pointer motion from the pointer of the seat
	 * associated with this object.
	 *
	 * A relative motion is in the same dimension as regular wl_pointer
	 * motion events, except they do not represent an absolute
	 * position. For example, moving a pointer from (x, y) to (x', y')
	 * would have the equivalent relative motion (x' - x, y' - y). If a
	 * pointer motion caused the absolute pointer position to be
	 * clipped by for example the edge of the monitor, the relative
	 * motion is unaffected by the clipping and will represent the
	 * unclipped motion.
	 *
	 * This event also contains non-accelerated motion deltas. The
	 * non-accelerated delta is, when applicable, the regular pointer
	 * motion delta as it was before having applied motion acceleration
	 * and other transformations such as normalization.
	 *
	 * Note that the non-accelerated delta does not represent 'raw'
	 * events as they were read from some device. Pointer motion
	 * acceleration is device- and configuration-specific and
	 * non-accelerated deltas and accelerated deltas may have the same
	 * value on some devices.
	 *
	 * Relative motions are not coupled to wl_pointer.motion events,
	 * and can be sent in combination with such events, but also
	 * independently. There may also be scenarios where
	 * wl_pointer.motion is sent, but there is no relative motion. The
	 * order of an absolute and relative motion event originating from
	 * the same physical motion is not guaranteed.
	 *
	 * If the client needs button events or focus state, it can receive
	 * them from a wl_pointer object of the same seat that the
	 * wp_relative_pointer object is associated with.
	 * @param utime_hi high 32 bits of a 64 bit timestamp with microsecond granularity
	 * @param utime_lo low 32 bits of a 64 bit timestamp with microsecond granularity
	 * @param dx the x component of the motion vector
	 * @param dy the y component of the motion vector
	 * @param dx_unaccel the x component of the unaccelerated motion vector
	 * @param dy_unaccel the y component of the unaccelerated motion vector
	 */
	void (*relative_motion)(void *data,
				struct zwp_relative_pointer_v1 *zwp_relative_pointer_v1,
				uint32_t utime_hi,
				uint32_t utime_lo,
				wl_fixed_t dx,
				wl_fixed_t dy,
				wl_fixed_t dx_unaccel,
				wl_fixed_t dy_unaccel);
};

/**
 * @ingroup iface_zwp_relative_pointer_v1
 */
static inline int
zwp_relative_pointer_v1_add_listener(struct zwp_relative_pointer_v1 *zwp_relative_pointer_v1,
				     const struct zwp_relative_pointer_v1_listener *listener, void *data)
{
	return wl_proxy_add_listener((struct wl_proxy *) zwp_relative_pointer_v1,
				     (void (**)(void)) listener, data);
}

#define ZWP_RELATIVE_POINTER_V1_DESTROY 0

/**
 * @ingroup iface_zwp_relative_pointer_v1
 */
#define ZWP_RELATIVE_POINTER_V1_RELATIVE_MOTION_SINCE_VERSION 1

/**
 * @ingroup iface_zwp_relative_pointer_v1
 */
#define ZWP_RELATIVE_POINTER_V1_DESTROY_SINCE_VERSION 1

/** @ingroup iface_zwp_relative_pointer_v1 */
static inline void
zwp_relative_pointer_v1_set_user_data(struct zwp_relative_pointer_v1 *zwp_relative_pointer_v1, void *user_data)
{
	wl_proxy_set_user_data((struct wl_proxy *) zwp_relative_pointer_v1, user_data);
}

/** @ingroup iface_zwp_relative_pointer_v1 */
static inline void *
zwp_relative_pointer_v1_get_user_data(struct zwp_relative_pointer_v1 *zwp_relative_pointer_v1)
{
	return wl_proxy_get_user_data((struct wl_proxy *) zwp_relative_pointer_v1);
}

static inline uint32_t
zwp_relative_pointer_v1_get_version(struct zwp_relative_pointer_v1 *zwp_relative_pointer_v1)
{
	return wl_proxy_get_version((struct wl_proxy *) zwp_relative_pointer_v1);
}

/**
 * @ingroup iface_zwp_relative_pointer_v1
 */
static inline void
zwp_relative_pointer_v1_destroy(struct zwp_relative_pointer_v1 *zwp_relative_pointer_v1)
{
	wl_proxy_marshal((struct wl_proxy *) zwp_relative_pointer_v1,
			 ZWP_RELATIVE_POINTER_V1_DESTROY);

	wl_proxy_destroy((struct wl_proxy *) zwp_relative_pointer_v1);
}

#ifdef  __cplusplus
}
#endif

#endif
