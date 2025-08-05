# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make cyclic_groups.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

cyclic_groups.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > cyclic_groups.tex

cyclic_groups.pdf: ../../*.tex ../../trademark.pdf *.tex cyclic_groups.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode cyclic_groups.tex
	../../latexrun cyclic_groups.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cyclic_groups.pdf
	open cyclic_groups.pdf

o:
	make open

terms: cyclic_groups.tex
	bbk sheet -mode ts cyclic_groups.tex

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
	rm -f cyclic_groups.*
	make reset
	bbk sheet -mode mk > Makefile