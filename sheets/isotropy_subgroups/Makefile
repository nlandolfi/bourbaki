# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make isotropy_subgroups.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

isotropy_subgroups.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > isotropy_subgroups.tex

isotropy_subgroups.pdf: ../../*.tex ../../trademark.pdf *.tex isotropy_subgroups.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode isotropy_subgroups.tex
	../../latexrun isotropy_subgroups.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: isotropy_subgroups.pdf
	open isotropy_subgroups.pdf

o:
	make open

terms: isotropy_subgroups.tex
	bbk sheet -mode ts isotropy_subgroups.tex

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
	rm -f isotropy_subgroups.*
	make reset
	bbk sheet -mode mk > Makefile