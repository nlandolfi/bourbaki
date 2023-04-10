# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make plane_vector_angles.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

plane_vector_angles.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > plane_vector_angles.tex

plane_vector_angles.pdf: ../../*.tex ../../trademark.pdf *.tex plane_vector_angles.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode plane_vector_angles.tex
	../../latexrun plane_vector_angles.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: plane_vector_angles.pdf
	open plane_vector_angles.pdf

o:
	make open

terms: plane_vector_angles.tex
	bbk sheet -mode ts plane_vector_angles.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f plane_vector_angles.*
	make reset
	bbk sheet -mode mk > Makefile