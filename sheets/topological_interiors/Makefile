# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make topological_interiors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

topological_interiors.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > topological_interiors.tex

topological_interiors.pdf: ../../*.tex ../../trademark.pdf *.tex topological_interiors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode topological_interiors.tex
	../../latexrun topological_interiors.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: topological_interiors.pdf
	open topological_interiors.pdf

o:
	make open

terms: topological_interiors.tex
	bbk sheet -mode ts topological_interiors.tex

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
	rm -f topological_interiors.*
	make reset
	bbk sheet -mode mk > Makefile