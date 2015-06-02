precision mediump float;

varying highp vec2 vUV;
varying highp vec3 vNormal;

uniform sampler2D uTexture;

const vec3 source_ambient_color=vec3(1.,1.,1.);
const vec3 source_diffuse_color=vec3(8.,8.,6.);
const vec3 source_specular_color=vec3(1.,1.,1.);
const vec3 source_direction = vec3(-1.0, 1.0, 1.0);

const vec3 mat_ambient_color=vec3(0.3,0.3,0.3);
const vec3 mat_diffuse_color=vec3(1.,1.,1.);
const vec3 mat_specular_color=vec3(1.,1.,1.);
const float mat_shininess=10.;

void main(void) {
		 vec3 dir = normalize(source_direction);
		 vec3 color = vec3(1.0);//vec3(texture2D(uTexture, vUV));
		 vec3 I_ambient = source_ambient_color * mat_ambient_color;
		 vec3 I_diffuse = source_diffuse_color * mat_diffuse_color * max(0.0, dot(vNormal, dir));
		 vec3 I = I_ambient + I_diffuse;
		 gl_FragColor = vec4(I * color, 1.0);
}
