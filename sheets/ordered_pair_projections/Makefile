# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make ordered_pair_projections.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

ordered_pair_projections.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > ordered_pair_projections.tex

ordered_pair_projections.pdf: ../../*.tex ../../trademark.pdf *.tex ordered_pair_projections.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode ordered_pair_projections.tex
	../../latexrun ordered_pair_projections.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: ordered_pair_projections.pdf
	open ordered_pair_projections.pdf

o:
	make open

terms: ordered_pair_projections.tex
	bbk sheet -mode ts ordered_pair_projections.tex

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
	rm -f ordered_pair_projections.*
	make reset
	bbk sheet -mode mk > Makefile