# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make conditional_dependency_graph.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

conditional_dependency_graph.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > conditional_dependency_graph.tex

conditional_dependency_graph.pdf: ../../*.tex ../../trademark.pdf *.tex conditional_dependency_graph.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode conditional_dependency_graph.tex
	../../latexrun conditional_dependency_graph.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: conditional_dependency_graph.pdf
	open conditional_dependency_graph.pdf

o:
	make open

terms: conditional_dependency_graph.tex
	bbk sheet -mode ts conditional_dependency_graph.tex

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
	rm -f conditional_dependency_graph.*
	make reset
	bbk sheet -mode mk > Makefile