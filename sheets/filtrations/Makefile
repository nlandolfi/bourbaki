# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make filtrations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

filtrations.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > filtrations.tex

filtrations.pdf: ../../*.tex ../../trademark.pdf *.tex filtrations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode filtrations.tex
	../../latexrun filtrations.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: filtrations.pdf
	open filtrations.pdf

o:
	make open

terms: filtrations.tex
	bbk sheet -mode ts filtrations.tex

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
	rm -f filtrations.*
	make reset
	bbk sheet -mode mk > Makefile