# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make supremum_norm_complete.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

supremum_norm_complete.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > supremum_norm_complete.tex

supremum_norm_complete.pdf: ../../*.tex ../../trademark.pdf *.tex supremum_norm_complete.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode supremum_norm_complete.tex
	../../latexrun supremum_norm_complete.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: supremum_norm_complete.pdf
	open supremum_norm_complete.pdf

o:
	make open

terms: supremum_norm_complete.tex
	bbk sheet -mode ts supremum_norm_complete.tex

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
	rm -f supremum_norm_complete.*
	make reset
	bbk sheet -mode mk > Makefile