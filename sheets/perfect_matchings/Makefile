# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make perfect_matchings.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

perfect_matchings.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > perfect_matchings.tex

perfect_matchings.pdf: ../../*.tex ../../trademark.pdf *.tex perfect_matchings.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode perfect_matchings.tex
	../../latexrun perfect_matchings.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: perfect_matchings.pdf
	open perfect_matchings.pdf

o:
	make open

terms: perfect_matchings.tex
	bbk sheet -mode ts perfect_matchings.tex

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
	rm -f perfect_matchings.*
	make reset
	bbk sheet -mode mk > Makefile