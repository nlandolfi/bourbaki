# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make greatest_lower_bounds.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

greatest_lower_bounds.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > greatest_lower_bounds.tex

greatest_lower_bounds.pdf: ../../*.tex ../../trademark.pdf *.tex greatest_lower_bounds.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode greatest_lower_bounds.tex
	../../latexrun greatest_lower_bounds.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: greatest_lower_bounds.pdf
	open greatest_lower_bounds.pdf

o:
	make open

terms: greatest_lower_bounds.tex
	bbk sheet -mode ts greatest_lower_bounds.tex

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
	rm -f greatest_lower_bounds.*
	make reset
	bbk sheet -mode mk > Makefile