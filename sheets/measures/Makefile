# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make measures.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

measures.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > measures.tex

measures.pdf: ../../*.tex ../../trademark.pdf *.tex measures.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode measures.tex
	../../latexrun measures.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: measures.pdf
	open measures.pdf

o:
	make open

terms: measures.tex
	bbk sheet -mode ts measures.tex

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
	rm -f measures.*
	make reset
	bbk sheet -mode mk > Makefile