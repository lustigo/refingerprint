package data

import "reflect"

// Extensions holds all possible WebGL extensions
// Source: https://www.khronos.org/registry/webgl/extensions/
var extensions = map[string]string{
	"NOT_FOUND":                             "WebGLExtMore",
	"OES_texture_float":                     "WebGLExtOESTextFloat",
	"OES_texture_half_float":                "WebGLExtOESTextHalfFloat",
	"WEBGL_lose_context":                    "WebGLExtLose",
	"WEBKIT_WEBGL_lose_context":             "WebGLExtLose",
	"OES_standard_derivatives":              "WebGLExtOESStd",
	"OES_vertex_array_object":               "WebGLExtOESVert",
	"WEBGL_debug_renderer_info":             "WebGLExtDebugRender",
	"WEBGL_debug_shaders":                   "WebGLExtDebugShader",
	"WEBGL_compressed_texture_s3tc":         "WebGLExtCompressTextS3TC",
	"WEBKIT_WEBGL_compressed_texture_s3tc":  "WebGLExtCompressTextS3TC",
	"WEBGL_depth_texture":                   "WebGLExtDepth",
	"WEBKIT_WEBGL_depth_texture":            "WebGLExtDepth",
	"OES_element_index_uint":                "WebGLExtOESEle",
	"EXT_texture_filter_anisotropic":        "WebGLExtTextFilter",
	"WEBKIT_EXT_texture_filter_anisotropic": "WebGLExtTextFilter",
	"EXT_frag_depth":                        "WebGLExtFrag",
	"WEBGL_draw_buffers":                    "WebGLExtDraw",
	"ANGLE_instanced_arrays":                "WebGLExtAngle",
	"OES_texture_float_linear":              "WebGLExtOESTextFloatLin",
	"OES_texture_half_float_linear":         "WebGLExtOESTextHalfFloatLin",
	"EXT_blend_minmax":                      "WebGLExtBlend",
	"EXT_shader_texture_lod":                "WebGLExtShader",
	"WEBGL_compressed_texture_pvrtc":        "WebGLExtCompressTextPVRTC",
	"EXT_color_buffer_half_float":           "WebGLExtColorBuffHalfFloat",
	"WEBGL_color_buffer_float":              "WebGLExtGLColorBuffFloat",
	"EXT_sRGB":                              "WebGLExtSrgb",
	"WEBGL_compressed_texture_etc1":         "WebGLExtCompressTextETC1",
	"EXT_disjoint_timer_query":              "WebGLExtDisjoint",
	"OES_fbo_render_mipmap":                 "WebGLExtOESFbo",
	"WEBGL_compressed_texture_etc":          "WebGLExtCompressTextETC",
	"WEBGL_compressed_texture_astc":         "WebGLExtCompressTextASTC",
	"EXT_color_buffer_float":                "WebGLExtColorBuffFloat",
	"WEBGL_compressed_texture_s3tc_srgb":    "WebGLExtCompressTextS3TCSrgb",
	"EXT_disjoint_timer_query_webgl2":       "WebGLExtDisjointGL2",
	"EXT_float_blend":                       "WebGLExtFloat",
	"OVR_multiview2":                        "WebGLExtOVR",
	"KHR_parallel_shader_compile":           "WebGLExtKHR",
	"EXT_texture_compression_bptc":          "WebGLExtTextCompressBPTC",
	"EXT_texture_compression_rgtc":          "WebGLExtTextCompressRGTC",
	"WEBGL_compressed_texture_atc":          "WebGLExtCompressTextATC", // Source: https://developer.mozilla.org/de/docs/Web/API/WebGL_API/Using_Extensions
}

// ExtractWebGLExtensions extracts the supported WebGL Extensions saves them
func (features *ProcessedFeatures) ExtractWebGLExtensions(winfo WebGLInfo) {
	f := reflect.ValueOf(features).Elem()
	for _, ext := range winfo.Extensions {
		v, o := extensions[ext]
		if !o {
			features.WebGLExtMore++
			continue
		}
		f.FieldByName(v).SetUint(1)
	}
}

// TODO: Graphiccard Vendor & Model
