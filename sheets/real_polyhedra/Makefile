# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_polyhedra.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

real_polyhedra.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > real_polyhedra.tex

real_polyhedra.pdf: ../../*.tex ../../trademark.pdf *.tex real_polyhedra.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_polyhedra.tex
	../../latexrun real_polyhedra.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_polyhedra.pdf
	open real_polyhedra.pdf

o:
	make open

terms: real_polyhedra.tex
	bbk sheet -mode ts real_polyhedra.tex

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
	rm -f real_polyhedra.*
	make reset
	bbk sheet -mode mk > Makefile