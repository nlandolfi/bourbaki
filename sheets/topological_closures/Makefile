# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make topological_closures.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

topological_closures.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > topological_closures.tex

topological_closures.pdf: ../../*.tex ../../trademark.pdf *.tex topological_closures.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode topological_closures.tex
	../../latexrun topological_closures.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: topological_closures.pdf
	open topological_closures.pdf

o:
	make open

terms: topological_closures.tex
	bbk sheet -mode ts topological_closures.tex

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
	rm -f topological_closures.*
	make reset
	bbk sheet -mode mk > Makefile